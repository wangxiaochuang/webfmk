package kernel

import (
	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/contract"
	"github.com/wxc/webfmk/framework/gin"
)

type FmkKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *FmkKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewFmkKernelService
}

func (provider *FmkKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *FmkKernelProvider) IsDefer() bool {
	return false
}

func (provider *FmkKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *FmkKernelProvider) Name() string {
	return contract.KernelKey
}
