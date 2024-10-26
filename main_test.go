package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
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
