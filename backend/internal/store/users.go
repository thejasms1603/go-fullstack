package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type UserStore struct {
	db *sql.DB
}


func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.Created, &user.Updated)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}