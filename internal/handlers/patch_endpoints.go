package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/t67y110v/dns-server/internal/models"
)

func (h *Handlers) PatchEndpoints(w http.ResponseWriter, req *http.Request) {

	var enpoint models.Endpoint

	err := json.NewDecoder(req.Body).Decode(&enpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(enpoint)
}
