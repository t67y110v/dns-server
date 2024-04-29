package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t67y110v/dns-server/internal/endpoints"
	"github.com/t67y110v/dns-server/internal/services"
)

func (h *Handlers) GetEndpoints(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("new request with queryy " + req.URL.RawQuery)
	target := chi.URLParam(req, "target")

	services := services.GetServices()
	for _, s := range services {
		if s == target {
			configuration := endpoints.GetServiceConfiguration(target)
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
