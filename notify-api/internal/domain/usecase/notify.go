package usecase

import (
	"context"
	"errors"

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

	res, err := uc.repository.FetchUser(ctx, n.UserID)
	if err != nil {
		return nil, err
	}

	if res.Notification {
		c, err := uc.repository.CreateNotify(ctx, n)
		if err != nil {
			return nil, err
		}

		return c, nil
	} else {
		return nil, errors.New("user not found or disabled notifications")
	}
}

func (uc *useCaseNotify) FetchNotify(ctx context.Context) (*entity.NotificationList, error) {
	res, err := uc.repository.FetchNotify(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *useCaseNotify) SendMessage(ctx context.Context, m *entity.Message) (*entity.Message, error) {
	res, err := uc.repository.SendMessage(ctx, m)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *useCaseNotify) FetchMessage(ctx context.Context) (*entity.MessageList, error) {
	res, err := uc.repository.FetchMessage(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
