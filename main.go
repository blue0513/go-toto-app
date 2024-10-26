package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	cfg, err := config.New()

	if err != nil {
		return err
	}

	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("Error: %+v\n", err)
	}

	return eg.Wait()
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
}
