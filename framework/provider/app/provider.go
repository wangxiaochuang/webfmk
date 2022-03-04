package app

import (
	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/contract"
)

type FmkAppProvider struct {
	BaseFolder string
}

func (f *FmkAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewFmkApp
}

func (f *FmkAppProvider) Boot(container framework.Container) error {
	return nil
}

func (f *FmkAppProvider) IsDefer() bool { return false }

func (f *FmkAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, f.BaseFolder}
}

func (f *FmkAppProvider) Name() string {
	return contract.AppKey
}
