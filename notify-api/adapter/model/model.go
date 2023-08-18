package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationRequest struct {
	DateTime time.Time `gorm:"not null"`
	Message  string    `gorm:"not null"`
}

type NotificationResponse struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"not null"`
	DateTime  time.Time `gorm:"not null"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}
