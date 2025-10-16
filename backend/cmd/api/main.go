package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/thejasms1603/go-fullstack/backend/internal/env"
)

	func main() {
		err:= godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
		port := env.GetString("PORT", ":8080")
		cfg:= Config{
			addr:port,
		}
		app:= &application{config: cfg}
		httpMux := app.mountRoutes()
		log.Fatal(app.run(httpMux))
	}