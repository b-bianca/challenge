package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	CPF string `json:"cpf" binding:"required"`
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	CPF          string    `json:"cpf"`
	Notification bool      `json:"notification"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OptoutRequest struct {
	Notification bool `json:"notification"`
}
