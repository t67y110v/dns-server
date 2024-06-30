package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t67y110v/dns-server/internal/models"
	"github.com/t67y110v/dns-server/internal/usecase"
)

func (h *Handlers) PatchEndpoints(w http.ResponseWriter, req *http.Request) {
	target := chi.URLParam(req, "target")
	var service models.Service

	err := json.NewDecoder(req.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = usecase.UpdateEndpoints(target, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
