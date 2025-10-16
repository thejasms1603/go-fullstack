package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)


type application struct {
	config Config
}


type Config struct {
	addr string
}

func (app *application) mountRoutes() *chi.Mux {
	r := chi.NewRouter()
	// Recovery middleware
	r.Use(middleware.Recoverer)
	// Logging middleware
	r.Use(middleware.Logger)
	// Request ID middleware
	r.Use(middleware.RequestID)
	// Real IP middleware
	r.Use(middleware.RealIP)
	// Timeout middleware
	r.Use(middleware.Timeout(60 * time.Second))
	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthCheckHandler)
	});
	return r
}

func (app * application) run(httpMux *chi.Mux) error {

	srv := &http.Server{
		Addr: app.config.addr,
		Handler: httpMux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: time.Minute,
	}
	log.Println("Server started on", app.config.addr, "...")
	if err:= srv.ListenAndServe(); err != nil {
		log.Println("Server failed to start: ", err)
		return err
	}
	return nil
	
}