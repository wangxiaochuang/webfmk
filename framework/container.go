package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	Bind(provider ServiceProvider) error
	IsBind(key string) bool

	Make(key string) (interface{}, error)
	MustMake(key string) interface{}
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type FmkContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]interface{}
	lock      sync.RWMutex
}

func NewFmkContainer() *FmkContainer {
	return &FmkContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (fmk *FmkContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range fmk.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (fmk *FmkContainer) Bind(provider ServiceProvider) error {
	fmk.lock.Lock()
	defer fmk.lock.Unlock()
	key := provider.Name()

	fmk.providers[key] = provider

	if provider.IsDefer() == false {
		instance, err := fmk.newInstance(provider, nil)
		if err != nil {
			return err
		}
		fmk.instances[key] = instance
	}
	return nil
}

func (fmk *FmkContainer) IsBind(key string) bool {
	return fmk.findServiceProvider(key) != nil
}

func (fmk *FmkContainer) findServiceProvider(key string) ServiceProvider {
	fmk.lock.RLock()
	defer fmk.lock.RUnlock()
	if sp, ok := fmk.providers[key]; ok {
		return sp
	}
	return nil
}

func (fmk *FmkContainer) Make(key string) (interface{}, error) {
	return fmk.make(key, nil, false)
}

func (fmk *FmkContainer) MustMake(key string) interface{} {
	serv, err := fmk.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (fmk *FmkContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return fmk.make(key, params, true)
}

func (fmk *FmkContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(fmk); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(fmk)
	}
	method := sp.Register(fmk)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (fmk *FmkContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	fmk.lock.RLock()
	defer fmk.lock.RUnlock()
	sp := fmk.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	if forceNew {
		return fmk.newInstance(sp, params)
	}

	if ins, ok := fmk.instances[key]; ok {
		return ins, nil
	}

	inst, err := fmk.newInstance(sp, params)
	if err != nil {
		return nil, err
	}

	fmk.instances[key] = inst
	return inst, nil
}
