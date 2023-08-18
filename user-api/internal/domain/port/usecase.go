package port

import (
	"context"

	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, u *entity.User) (*entity.User, error)
	PartialUpdateUser(ctx context.Context, u *entity.User) error
}
