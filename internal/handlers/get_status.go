package handlers

import "net/http"

func (h *Handlers) GetStatus(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
