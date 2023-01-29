package core

import "net/http"

type handlers struct {
}

func (h *handlers) find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJson(w, http.StatusOK, response)
}
