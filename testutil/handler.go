package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg interface{}
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Errorf("Error: %+v\n", err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Errorf("Error: %+v\n", err)
	}
	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Errorf("(-want +got):\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })

	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("watnt status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}
	if len(gb) == 0 && len(body) == 0 {
		return
	}
	AssertJSON(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Error: %+v\n", err)
	}
	return bt
}
