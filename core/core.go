package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Core struct {
	edpoint string
	Mux     *mux.Router
}

func New(endpoint string, setup setupRoute, Mux *mux.Router) Core {
	core := Core{edpoint: endpoint}

	core.setup(setup)

	return core

}

func (core *Core) setup(setup setupRoute) {

	endpoint := "/" + core.edpoint

	if setup.find {
		core.Mux.HandleFunc(endpoint, func(response http.ResponseWriter, response *http.Response) {

		})
	}
}
