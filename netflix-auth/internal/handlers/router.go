package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type router struct {
	*mux.Router
	chainFunc func(hf http.HandlerFunc) http.Handler
}

func newRouter() *router {
	return &router{
		Router: mux.NewRouter(),
		chainFunc: func(hf http.HandlerFunc) http.Handler {
			return hf
		},
	}
}

func (r *router) chain(c alice.Chain) *router {
	r.chainFunc = c.ThenFunc

	return r
}

func (r *router) subRouter(prefix string) *router {
	return &router{
		Router:    r.PathPrefix(prefix).Subrouter(),
		chainFunc: r.chainFunc,
	}
}

func (r *router) handle(pattern string, handlerFunc http.HandlerFunc, method string) {
	r.Handle(pattern, r.chainFunc(handlerFunc)).Methods(method)
}
