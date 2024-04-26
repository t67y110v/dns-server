package handlers

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEndpoints(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/endpoints?target=hello_world_server", nil)
	w := httptest.NewRecorder()
	h := NewHandlers(slog.Default())
	h.GetEndpoints(w, req)
	res := w.Result()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	want := "200 OK"
	got := res.Status
	if got != want {
		t.Error("wrong status code, wanted " + want + " actual " + got)
	}

	if data == nil {
		t.Error("empty response ")
	}
}
