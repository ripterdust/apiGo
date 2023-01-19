package router

import (
	"github.com/gorilla/mux"
	"prueba.com/core"
	"prueba.com/users"
)

type Router struct {
	routes []func(*mux.Router) core.Core
}

func New() {

	routes := []func(*mux.Router) core.Core{
		users.UsersRoute,
	}

	router := Router{
		routes: routes,
	}

	router.setRoutes()

}

func (router *Router) setRoutes() {
	Mux := mux.NewRouter()

	for _, route := range router.routes {
		route(Mux)
	}

}
