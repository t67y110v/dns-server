package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/t67y110v/dns-server/internal/endpoints"
)

func (h *Handlers) GetEndpoints(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("new request with queryy " + req.URL.RawQuery)

	target := req.URL.Query().Get("target")

	resp := endpoints.GetResponse()
	switch target {
	case "hello_world_server":
		marshalledResponse, err := json.Marshal(resp)
		if err != nil {
			h.logger.Error("unmarshall response error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(marshalledResponse)
		if err != nil {
			h.logger.Error("writing response error", err)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
