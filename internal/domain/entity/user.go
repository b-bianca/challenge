package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name         string    `gorm:"not null"`
	CPF          string    `gorm:"not null"`
	Email        string    `gorm:"not null"`
	Password     string    `gorm:"not null"`
	Notification bool      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime"`
}
