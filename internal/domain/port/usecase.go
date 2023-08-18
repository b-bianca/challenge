package port

import (
	"context"

	"github.com/b-bianca/melichallenge/internal/domain/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, u *entity.User) (*entity.User, error)
	OptoutUser(ctx context.Context, u *entity.User) error
}
