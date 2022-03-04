package kernel

import (
	"net/http"

	"github.com/wxc/webfmk/framework/gin"
)

type FmkKernelService struct {
	engine *gin.Engine
}

func NewFmkKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &FmkKernelService{engine: httpEngine}, nil
}

func (s *FmkKernelService) HttpEngine() http.Handler {
	return s.engine
}
