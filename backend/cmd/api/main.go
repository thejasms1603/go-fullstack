package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/thejasms1603/go-fullstack/backend/internal/db"
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
			db : dbConfig{
				addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/go-fullstack?sslmode=disable"),
				maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
				maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
				maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "60s"),
			},
		}
		db, err := db.NewConnectionPool(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
		if err != nil {
			log.Fatal("Failed to create database connection: ", err)
		}
		defer db.Close()
		log.Println("Database connection pool created successfully")

		store:= store.NewStorage(db)

		app:= &application{config: cfg, store: store}

		httpMux := app.mountRoutes()

		log.Fatal(app.run(httpMux))
	}