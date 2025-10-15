package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	api:= &apiConfig{addr: ":8080"}

	router:= chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	
	v1Router:= chi.NewRouter()
	v1Router.Get("/users", api.getUsersHandler)
	v1Router.Post("/users", api.createUserHandler)
	router.Mount("/v1", v1Router)
	
	srv := &http.Server{
		Addr: api.addr,
		Handler: router,
	}

	
	log.Println("Server started on ", api.addr)
	if err:= srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}