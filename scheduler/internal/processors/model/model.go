package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	DateTime  time.Time `json:"date_time"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotificationList struct {
	Result []*Notification
}

type MessageRequest struct {
	NotifyID uuid.UUID `json:"notify_id"`
	Message  string    `json:"message"`
}
