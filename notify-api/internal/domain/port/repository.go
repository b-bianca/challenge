package port

import (
	"context"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
)

type NotifyRepository interface {
	CreateNotify(ctx context.Context, u *entity.Notification) (*entity.Notification, error)
	FetchNotify(ctx context.Context) (*entity.NotificationList, error)
}
