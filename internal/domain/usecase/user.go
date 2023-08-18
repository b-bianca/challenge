package usecase

import (
	"context"

	"github.com/b-bianca/melichallenge/internal/domain/entity"
	"github.com/b-bianca/melichallenge/internal/domain/port"
)

type useCaseUser struct {
	repository port.UserRepository
}

func NewCustomerUseCase(user port.UserRepository) port.UserUseCase {
	return &useCaseUser{
		repository: user,
	}
}

func (uc *useCaseUser) CreateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	res, err := uc.repository.CreateUser(ctx, u)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *useCaseUser) OptoutUser(ctx context.Context, u *entity.User) error {
	err := uc.repository.OptoutUser(ctx, u)

	if err != nil {
		return err
	}

	return nil
}
