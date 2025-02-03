package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Define routes for the application
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// CORS Middleware
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from React frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middleware
	mux.Use(middleware.Heartbeat("/ping")) // Health check endpoint
	mux.Use(middleware.Recoverer)          // Recover from panics
	mux.Use(middleware.Logger)             // Log requests

	// Calls LogHandler for /log requests
	mux.Post("/log", app.LogHandler)

	// Calls GetLogsHandler for /logs requests (GET request to fetch logs by service name)
	mux.Get("/logs", app.GetLogsHandler)

	mux.Get("/logs/all", app.GetAllLogsHandler)

	return mux
}
