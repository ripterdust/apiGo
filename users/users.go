package users

import (
	"net/http"

	"prueba.com/core"
)

func UsersRoute(Mux *http.ServeMux) core.Core {

	routeSetup := core.SetupRoute{Find: true}

	route := core.New("users", routeSetup, Mux)

	return route

}
