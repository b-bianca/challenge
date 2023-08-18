package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	CPF      string `json:"cpf" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CPF          string    `json:"cpf"`
	Email        string    `json:"email"`
	Notification bool      `json:"notification"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OptoutRequest struct {
	Notification bool `json:"notification"`
}
