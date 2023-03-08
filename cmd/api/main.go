package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{locale}", handleHome)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
