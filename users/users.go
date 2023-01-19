package users

import (
	"github.com/gorilla/mux"
	"prueba.com/core"
)

func UsersRoute(Mux *mux.Router) core.Core {

	routeSetup := core.NewSetup(true, true, true, true)

	route := core.New("usuario", routeSetup, Mux)

	return route

}
