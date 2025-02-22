package main

import (
	"go_todo_app/handler"
	"go_todo_app/store"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Contest-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": ok}`))
	})

	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux
}
