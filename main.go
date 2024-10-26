package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		return err
	}

	handler := NewMux()
	s := NewServer(l, handler)
	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
}
