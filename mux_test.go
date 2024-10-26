package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	mux := NewMux()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error: %+v\n", err)
	}

	want := `{"status": ok}`
	if string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
