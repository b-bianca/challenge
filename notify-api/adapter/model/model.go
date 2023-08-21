package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationRequest struct {
	DateTime time.Time `json:"date_time" binding:"required"`
	Message  string    `json:"message" binding:"required"`
}

type NotificationResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	DateTime  time.Time `json:"date_time"`
	Message   string    `json:"message"`
	Ack       bool      `json:"ack"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotificationListResponse struct {
	Result []*NotificationResponse
}

type MessageRequest struct {
	NotifyID uuid.UUID `json:"notify_id"`
	Message  string    `json:"message"`
}
type MessageResponse struct {
	ID        uuid.UUID `json:"id"`
	NotifyID  uuid.UUID `json:"notify_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"updated_at"`
}

type MessageListResponse struct {
	Result []*MessageResponse
}
