package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func initRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(LogRequest, SecureHeaders, CorsPolicy, RateLimitter)
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return router
}
