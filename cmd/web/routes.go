package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters",handlers.Repo.Generals)
	mux.Get("/majors-suite",handlers.Repo.Majors)
	mux.Get("/make-reservation",handlers.Repo.Reservation)
	mux.Get("/search-availability",handlers.Repo.Availability)
	mux.Post("/search-availability",handlers.Repo.PostAvailability)
	mux.Get("/contact",handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}