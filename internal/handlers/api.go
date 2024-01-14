package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/kaus19/online_offline_tracker/internal/middleware"
)

func Handler(r *chi.Mux){
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		//Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/status", GetUserStatus)
	})

}