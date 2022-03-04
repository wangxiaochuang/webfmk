package contract

import "net/http"

const KernelKey = "fmk:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
