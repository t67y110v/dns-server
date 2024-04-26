package handlers

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()
	h := NewHandlers(slog.Default())
	h.GetStatus(w, req)
	res := w.Result()
	defer res.Body.Close()
	want := "200 OK"
	got := res.Status
	if got != want {
		t.Error("wrong status code, wanted " + want + " actual " + got)
	}
}
