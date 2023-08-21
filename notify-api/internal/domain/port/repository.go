package port

import (
	"context"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/google/uuid"
)

type NotifyRepository interface {
	CreateNotify(ctx context.Context, u *entity.Notification) (*entity.Notification, error)
	FetchNotify(ctx context.Context) (*entity.NotificationList, error)
	SendMessage(ctx context.Context, m *entity.Message) (*entity.Message, error)
	FetchMessage(ctx context.Context) (*entity.MessageList, error)
	FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}
