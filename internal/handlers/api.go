package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		InitiateDb()
		router.Get("/status", GetUserStatus)
		router.Get("/all", GetAllStatus)
	})

}
