package handlers

import "net/http"

func (h *Handlers) Status(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
