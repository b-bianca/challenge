package usecase

import (
	"context"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/port"
)

type useCaseNotify struct {
	repository port.NotifyRepository
}

func NewNotifyUseCase(nt port.NotifyRepository) port.NotifyUseCase {
	return &useCaseNotify{
		repository: nt,
	}
}

func (uc *useCaseNotify) CreateNotify(ctx context.Context, n *entity.Notification) (*entity.Notification, error) {
	res, err := uc.repository.CreateNotify(ctx, n)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *useCaseNotify) FetchNotify(ctx context.Context) (*entity.NotificationList, error) {
	res, err := uc.repository.FetchNotify(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
