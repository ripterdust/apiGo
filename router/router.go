package router

import (
	"net/http"

	"prueba.com/core"
	"prueba.com/users"
)

type Router struct {
	routes []func(*http.ServeMux) core.Core
	Mux    *http.ServeMux
}

func New(muxRouter *http.ServeMux) Router {

	routes := []func(*http.ServeMux) core.Core{
		users.UsersRoute,
	}

	router := Router{
		routes: routes,
		Mux:    muxRouter,
	}

	router.setRoutes()

	return router
}

func (router *Router) setRoutes() {

	for _, route := range router.routes {
		route(router.Mux)
	}

}
