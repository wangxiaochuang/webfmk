package http

import (
	"github.com/wxc/webfmk/app/http/module/demo"
	"github.com/wxc/webfmk/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
