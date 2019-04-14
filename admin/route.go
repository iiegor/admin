package admin

import (
	"net/http"
	"sync/atomic"

	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

func ChainMiddlewares(f httprouter.Handle, middlewares ...Middleware) httprouter.Handle {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func MetricsMiddleware(resource *Resource) Middleware {
	return func(f httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			atomic.AddInt64(&resource.metrics.hits, 1)

			f(w, r, p)
		}
	}
}

func AuthMiddleware() Middleware {
	return func(f httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			// ..

			f(w, r, p)
		}
	}
}
