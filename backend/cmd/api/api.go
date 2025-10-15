package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type apiConfig struct {
	addr string
}

var users = []User{}


func (cfg *apiConfig) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := ValidateUser(payload); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user := User{
		ID: uuid.New().String(),
		Name: payload.Name,
		Email: payload.Email,
		Created: time.Now(),
		Updated: time.Now(),
	}
	users = append(users, user)
	respondWithJSON(w, http.StatusCreated, user )
}

func ValidateUser(u User) error{
	if u.Name == "" || u.Email == "" {
		return errors.New("name and email are required")
	}
	for _, user := range users {
		if user.Email == u.Email || user.Name == u.Name {
			return errors.New("email and name must be unique")
		}
	}
	return nil
}