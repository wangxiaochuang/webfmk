package http

import (
	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/gin"
)

func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.SetContainer(container)
	Routes(r)

	return r, nil
}
