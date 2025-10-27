package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/thejasms1603/go-fullstack/backend/internal/env"
	"github.com/thejasms1603/go-fullstack/backend/internal/store"
)

	func main() {
		err:= godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
		cfg:= Config{
			addr: env.GetString("ADDR", ":8080"),
		}
		store:= store.NewStorage(nil)
		app:= &application{config: cfg, store: store}
		httpMux := app.mountRoutes()
		log.Fatal(app.run(httpMux))
	}