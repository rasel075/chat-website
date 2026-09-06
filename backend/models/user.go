package models

import "time"

// User represents a user account
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // "-" means don't include in JSON responses
	CreatedAt time.Time `json:"created_at"`
}