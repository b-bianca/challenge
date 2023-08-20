package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/port/mocks"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	defaultID  = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")
	ctxDefault = context.Background()
)

func Test_CreateUser(t *testing.T) {
	t.Run("when everything goes as expected; should return user and no error", func(t *testing.T) {

		repo := mocks.NewUserRepository(t)
		service := usecase.NewCustomerUseCase(repo)

		given := &entity.User{
			CPF: "12345678912",
		}

		want := &entity.User{
			ID:           defaultID,
			CPF:          "12345678912",
			Notification: true,
		}

		repo.On("CreateUser", ctxDefault, given).Return(want, nil).Once()

		got, err := service.CreateUser(ctxDefault, given)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewUserRepository(t)
		service := usecase.NewCustomerUseCase(repo)

		wantError := errors.New("error")

		given := &entity.User{
			CPF: "12345678912",
		}
		want := &entity.User{
			ID:           defaultID,
			CPF:          "12345678912",
			Notification: true,
		}
		repo.On("CreateUser", ctxDefault, given).Return(want, wantError).Once()

		got, err := service.CreateUser(ctxDefault, given)
		assert.ErrorIs(t, err, wantError)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}

func TestPartialUpdate(t *testing.T) {
	given := &entity.User{
		Notification: false,
	}
	t.Run("when everything goes as expected; should return no error", func(t *testing.T) {
		repo := mocks.NewUserRepository(t)
		service := usecase.NewCustomerUseCase(repo)

		repo.On("PartialUpdateUser", context.Background(), given).Return(nil).Once()

		err := service.PartialUpdateUser(context.Background(), given)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("when update returns error; should send forward the error", func(t *testing.T) {
		repo := mocks.NewUserRepository(t)
		service := usecase.NewCustomerUseCase(repo)

		wantedErr := errors.New("error")
		repo.On("PartialUpdateUser", context.Background(), given).Return(wantedErr).Once()

		err := service.PartialUpdateUser(context.Background(), given)
		assert.ErrorIs(t, err, wantedErr)
		repo.AssertExpectations(t)
	})
}
