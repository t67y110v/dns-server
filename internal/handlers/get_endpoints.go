package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t67y110v/dns-server/internal/services"
	"github.com/t67y110v/dns-server/internal/usecase"
)

func (h *Handlers) GetEndpoints(w http.ResponseWriter, req *http.Request) {
	target := chi.URLParam(req, "target")

	services := services.GetServiceNames()
	for _, s := range services {
		if s == target {
			configuration := usecase.GetServiceConfiguration(target)
			marshalledResponse, err := json.Marshal(configuration)
			if err != nil {
				h.logger.Error("unmarshall response error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(marshalledResponse)
			if err != nil {
				h.logger.Error("writing response error", err)
			}
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)

}
