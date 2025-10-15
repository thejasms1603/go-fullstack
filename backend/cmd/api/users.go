package main

import "time"

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}