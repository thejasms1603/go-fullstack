package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type apiConfig struct {
	addr string
}


func main() {
	api:= &apiConfig{
		addr: ":8080",
	}
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	
	// v1Router
	v1Router:= chi.NewRouter()
	v1Router.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"))
	})

	router.Mount("/v1", v1Router)

	srv:= &http.Server{
		Addr: api.addr,
		Handler: router,
	}
	log.Println("Server started on ", api.addr)
	if err:= srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}