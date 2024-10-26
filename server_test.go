package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Errorf("Error: %+v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	})

	eg.Go(func() error {
		s := NewServer(l, mux)
		return s.Run(ctx)
	})

	in := "message"
	rsp, err := http.Get("http://localhost:8080/" + in)
	if err != nil {
		t.Errorf("Error: %+v\n", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("Error: %+v\n", err)
	}

	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Errorf("Error: %+v\n", err)
	}
}
