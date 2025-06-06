package entity

import (
	"time"

	"github.com/google/uuid"
)

type AdminUser struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
}

type AdminCredentials struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
}

type Session struct {
	ID         uuid.UUID `json:"id"`         
	UserEmail  string    `json:"user_email"` 
	CreatedAt  time.Time `json:"created_at"` 
	ExpiresAt  time.Time `json:"expires_at"` 
}

type SessionCreationResp struct {
	ID         uuid.UUID `json:"id"` 
	ExpiresAt  time.Time `json:"expires_at"` 
}