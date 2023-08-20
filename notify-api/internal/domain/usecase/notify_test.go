package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/port/mocks"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	defaultID  = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")
	ctxDefault = context.Background()
)

func TestCreateNotify(t *testing.T) {
	t.Run("when everything goes as expected; should return notify and no error", func(t *testing.T) {

		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		given := &entity.Notification{
			DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			Message:  "message",
		}

		want := &entity.Notification{
			ID:       defaultID,
			UserID:   defaultID,
			DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			Message:  "message",
		}

		repo.On("CreateNotify", ctxDefault, given).Return(want, nil).Once()

		got, err := service.CreateNotify(ctxDefault, given)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wantError := errors.New("error")
		given := &entity.Notification{
			DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			Message:  "message",
		}
		want := &entity.Notification{
			ID:       defaultID,
			UserID:   defaultID,
			DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			Message:  "message",
		}
		repo.On("CreateNotify", ctxDefault, given).Return(want, wantError).Once()

		got, err := service.CreateNotify(ctxDefault, given)
		assert.ErrorIs(t, err, wantError)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}

func TestFetchNotify(t *testing.T) {
	t.Run("when everything goes as expected; should return notify and no error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wanted := &entity.NotificationList{
			Result: []*entity.Notification{
				{
					ID:       defaultID,
					UserID:   defaultID,
					DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
					Message:  "message",
				},
			},
		}

		repo.On("FetchNotify", context.Background()).Return(wanted, nil).Once()

		got, err := service.FetchNotify(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, got, wanted)
		repo.AssertExpectations(t)
	})
	t.Run("when repository returns error; should send forward the error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wantedErr := errors.New("error")
		repo.On("FetchNotify", context.Background()).Return(nil, wantedErr).Once()

		got, err := service.FetchNotify(context.Background())
		assert.ErrorIs(t, err, wantedErr)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("when everything goes as expected; should return message and no error", func(t *testing.T) {

		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		given := &entity.Message{
			NotifyID: defaultID,
			Message:  "message",
		}

		want := &entity.Message{
			ID:       defaultID,
			NotifyID: defaultID,
			Message:  "message",
		}

		repo.On("SendMessage", ctxDefault, given).Return(want, nil).Once()

		got, err := service.SendMessage(ctxDefault, given)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wantError := errors.New("error")
		given := &entity.Message{
			NotifyID: defaultID,
			Message:  "message",
		}

		repo.On("SendMessage", ctxDefault, given).Return(nil, wantError).Once()

		got, err := service.SendMessage(ctxDefault, given)
		assert.ErrorIs(t, err, wantError)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}

func TestFetchMessage(t *testing.T) {
	t.Run("when everything goes as expected; should return message list and no error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wanted := &entity.MessageList{
			Result: []*entity.Message{
				{
					ID:       defaultID,
					NotifyID: defaultID,
					Message:  "message",
				},
			},
		}

		repo.On("FetchMessage", context.Background()).Return(wanted, nil).Once()

		got, err := service.FetchMessage(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, got, wanted)
		repo.AssertExpectations(t)
	})
	t.Run("when repository returns error; should send forward the error", func(t *testing.T) {
		repo := mocks.NewNotifyRepository(t)
		service := usecase.NewNotifyUseCase(repo)

		wantedErr := errors.New("error")
		repo.On("FetchMessage", context.Background()).Return(nil, wantedErr).Once()

		got, err := service.FetchMessage(context.Background())
		assert.ErrorIs(t, err, wantedErr)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}
