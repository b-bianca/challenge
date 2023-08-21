package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"not null"`
	DateTime  time.Time `gorm:"not null"`
	Message   string    `gorm:"not null"`
	Ack       bool
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type NotificationList struct {
	Result []*Notification
}

type Message struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	NotifyID  uuid.UUID `gorm:"not null"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}

type MessageList struct {
	Result []*Message
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CPF          string    `gorm:"not null"`
	Notification bool      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime"`
}
