package core

import (
	"net/http"
)

type Core struct {
	edpoint string
	Mux     *http.ServeMux
}

func New(endpoint string, setup SetupRoute, Mux *http.ServeMux) Core {
	core := Core{edpoint: endpoint, Mux: Mux}

	core.setup(setup)

	return core

}

func (core *Core) setup(setup SetupRoute) {

	endpoint := "/" + core.edpoint
	handler := handlers{}

	if setup.Find {
		core.Mux.HandleFunc(endpoint, handler.find)
	}
}
