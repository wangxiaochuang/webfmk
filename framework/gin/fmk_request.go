package gin

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

type IRequest interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}
	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// json body
	BindJSON(obj interface{}) error

	// xml body
	BindXML(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	GetHeader(key string) string

	// cookie
	GetCookie(key string) (string, error)
}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	if val, exist := ctx.GetQuery(key); exist {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	if val, exist := ctx.GetQuery(key); exist {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	if val, exist := ctx.GetQuery(key); exist {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	if val, exist := ctx.GetQuery(key); exist {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	if val, exist := ctx.GetQuery(key); exist {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	return ctx.GetQuery(key)
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	return ctx.GetQueryArray(key)
}

func (ctx *Context) FmkParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := ctx.FmkParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParam(key string) interface{} {
	return ctx.FmkParam(key)
}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	if val, ok := ctx.GetPostForm(key); ok {
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	if values, ok := ctx.GetPostFormArray(key); ok {
		return values, ok
	}
	return def, false
}

func (ctx *Context) DefaultFormFile(name string) (*multipart.FileHeader, error) {
	return ctx.FormFile(name)
}

func (ctx *Context) DefaultForm(key string) interface{} {
	if values, ok := ctx.GetPostFormArray(key); ok {
		return values
	}
	return nil
}

func (ctx *Context) Uri(key string) string {
	return ctx.Request.URL.Path
}

func (ctx *Context) Method(key string) string {
	return ctx.Request.Method
}

func (ctx *Context) Host(key string) string {
	return ctx.Request.URL.Host
}

func (ctx *Context) ClientIp(key string) string {
	return ctx.Request.RemoteAddr
}

func (ctx *Context) GetCookie(key string) (string, error) {
	return ctx.Cookie(key)
}
