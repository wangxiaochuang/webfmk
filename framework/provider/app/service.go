package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/contract"
	"github.com/wxc/webfmk/framework/util"
)

type FmkApp struct {
	contract.App
	container  framework.Container
	baseFolder string
}

func (f FmkApp) Version() string {
	return "0.0.1"
}

func (f FmkApp) BaseFolder() string {
	if f.baseFolder != "" {
		return f.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder param, default to .")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
}

func (f FmkApp) ConfigFolder() string {
	return filepath.Join(f.BaseFolder(), "config")
}

func (f FmkApp) LogFolder() string {
	return filepath.Join(f.StorageFolder(), "log")
}

func (f FmkApp) HttpFolder() string {
	return filepath.Join(f.BaseFolder(), "http")
}

func (f FmkApp) ConsoleFolder() string {
	return filepath.Join(f.BaseFolder(), "console")
}

func (f FmkApp) StorageFolder() string {
	return filepath.Join(f.BaseFolder(), "storage")
}

func (f FmkApp) ProviderFolder() string {
	return filepath.Join(f.BaseFolder(), "provider")
}

func (f FmkApp) MiddlewareFolder() string {
	return filepath.Join(f.HttpFolder(), "middleware")
}

func (f FmkApp) CommandFolder() string {
	return filepath.Join(f.ConsoleFolder(), "command")
}

func (f FmkApp) RuntimeFolder() string {
	return filepath.Join(f.StorageFolder(), "runtime")
}

func (f FmkApp) TestFolder() string {
	return filepath.Join(f.BaseFolder(), "test")
}

func NewFmkApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &FmkApp{baseFolder: baseFolder, container: container}, nil
}
