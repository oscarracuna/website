package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/oscarracuna/website/pkg/config"
	"github.com/oscarracuna/website/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// Middleware must be defined before routes
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)

	return mux
}
