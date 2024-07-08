package handlers

import (
	"github.com/go-chi/chi"
	chimddle "github.com/go-chi/chi/v5/middleware"
	"github.com/swagatswaroop/goapi/internal/middleware"
)

// The function started with capital letter suggests that it can be used in other files
// if started with a lowercase its a private function and can be used in this package only
func Handler(r *chi.Mux) {
	// Global middlware
	r.Use(chimddle.StripSlashes) //this makes the trailing slashes should always be ignored that https://localhost:5000/api/ <- this trailing slashes
	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
