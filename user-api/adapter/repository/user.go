package repository

import (
	"context"

	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{db}
}

func (us *User) CreateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	dbFn := us.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("user").Create(u); result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (us *User) PartialUpdateUser(ctx context.Context, u *entity.User) error {
	dbFn := us.db.WithContext(ctx)

	return dbFn.Table("user").Where(u.ID).Updates(u).Error
}
