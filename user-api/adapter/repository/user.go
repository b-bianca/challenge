package repository

import (
	"context"
	"fmt"

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

	result := dbFn.Table("users").Create(u)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("CPF already exists")
	}

	return u, nil
}

func (us *User) PartialUpdateUser(ctx context.Context, u *entity.User) error {
	dbFn := us.db.WithContext(ctx)

	return dbFn.Table("users").Where(u.ID).Updates(u).Error
}
