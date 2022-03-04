package main

import (
	"github.com/wxc/webfmk/app/console"
	"github.com/wxc/webfmk/app/http"
	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/provider/app"
	"github.com/wxc/webfmk/framework/provider/kernel"
)

func main() {
	container := framework.NewFmkContainer()
	container.Bind(&app.FmkAppProvider{})

	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.FmkKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
