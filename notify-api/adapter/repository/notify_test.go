package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/b-bianca/melichallenge/notify-api/adapter/repository"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	createNotification = `
	INSERT INTO "notify" (.+)
	VALUES (.+)
	ON CONFLICT DO NOTHING RETURNING "id"
`

	fetchNotification = `SELECT (.+) FROM "notify"  WHERE (.+)`

	createMessage = `
	INSERT INTO "message" (.+)
	VALUES (.+)
	ON CONFLICT DO NOTHING RETURNING "id"
`

	fetchMessage = `SELECT (.+) FROM "message"`
)

var (
	defaultID = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")

	ctx = context.Background()

	notification = &entity.Notification{
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "message",
	}

	message = &entity.Message{
		NotifyID: defaultID,
		Message:  "message",
	}
)

func TestCreateNotify(t *testing.T) {
	t.Run("when everything goes ok, should create a notification register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		mock.ExpectBegin()
		mock.
			ExpectQuery(createNotification).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(defaultID))
		mock.ExpectCommit()

		r := repository.New(dbGorm)
		result, err := r.Notify.CreateNotify(ctx, notification)

		assert.NoError(t, err)
		assert.Equal(t, result.ID, defaultID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.ExpectBegin()
		mock.ExpectQuery(createNotification).WillReturnError(wantErr)
		mock.ExpectRollback()

		r := repository.New(dbGorm)
		result, err := r.Notify.CreateNotify(ctx, notification)
		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, result)
	})
}

func TestFetchNotify(t *testing.T) {
	t.Run("when everything goes ok, should return a notification list register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		wantRows := []*entity.Notification{
			{
				ID:       defaultID,
				UserID:   defaultID,
				Message:  "message",
				DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			},
			{
				ID:       defaultID,
				UserID:   defaultID,
				Message:  "message-test",
				DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
			},
		}

		rows := sqlmock.NewRows([]string{"id", "user_id", "message", "date_time"}).
			AddRow(
				wantRows[0].ID,
				wantRows[0].UserID,
				wantRows[0].Message,
				wantRows[0].DateTime,
			).
			AddRow(
				wantRows[1].ID,
				wantRows[1].UserID,
				wantRows[1].Message,
				wantRows[1].DateTime,
			)

		mock.
			ExpectQuery(fetchNotification).WillReturnRows(rows)

		r := repository.New(dbGorm)
		got, err := r.Notify.FetchNotify(ctx)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Equal(t, wantRows, got.Result)
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.
			ExpectQuery(fetchNotification).WillReturnError(wantErr)

		r := repository.New(dbGorm)
		got, err := r.Notify.FetchNotify(ctx)

		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, got)
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("when everything goes ok, should create a message register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		mock.ExpectBegin()
		mock.
			ExpectQuery(createMessage).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(defaultID))
		mock.ExpectCommit()

		r := repository.New(dbGorm)
		result, err := r.Notify.SendMessage(ctx, message)

		assert.NoError(t, err)
		assert.Equal(t, result.ID, defaultID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("when db returns unmapped error, should propagate an error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.ExpectBegin()
		mock.ExpectQuery(createMessage).WillReturnError(wantErr)
		mock.ExpectRollback()

		r := repository.New(dbGorm)
		result, err := r.Notify.SendMessage(ctx, message)
		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, result)
	})
}

func TestFetchMessage(t *testing.T) {
	t.Run("when everything goes ok, should return a message list register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		wantRows := []*entity.Message{
			{
				ID:       defaultID,
				NotifyID: defaultID,
				Message:  "message",
			},
			{
				ID:       defaultID,
				NotifyID: defaultID,
				Message:  "message-test",
			},
		}

		rows := sqlmock.NewRows([]string{"id", "notify_id", "message"}).
			AddRow(
				wantRows[0].ID,
				wantRows[0].NotifyID,
				wantRows[0].Message,
			).
			AddRow(
				wantRows[1].ID,
				wantRows[1].NotifyID,
				wantRows[1].Message,
			)

		mock.
			ExpectQuery(fetchMessage).WillReturnRows(rows)

		r := repository.New(dbGorm)
		got, err := r.Notify.FetchMessage(ctx)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Equal(t, wantRows, got.Result)
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.
			ExpectQuery(fetchMessage).WillReturnError(wantErr)

		r := repository.New(dbGorm)
		got, err := r.Notify.FetchMessage(ctx)

		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, got)
	})
}
