package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t67y110v/dns-server/internal/endpoints"
	"github.com/t67y110v/dns-server/internal/models"
)

func (h *Handlers) PostEndpoints(w http.ResponseWriter, req *http.Request) {
	target := chi.URLParam(req, "target")
	ep := models.Service{}

	err := json.NewDecoder(req.Body).Decode(&ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = endpoints.SetEndpoints(target, ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
