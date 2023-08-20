package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/b-bianca/melichallenge/user-api/adapter/repository"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	createUser = `
	INSERT INTO "users" (.+)
	VALUES (.+)
	ON CONFLICT DO NOTHING RETURNING "id"
`

	updatePartial = `UPDATE "users" SET (.+) WHERE (.+)`
)

var (
	defaultID = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")

	ctx = context.Background()

	user = &entity.User{
		CPF: "12345678933",
	}

	partialUpdated = &entity.User{
		Notification: false,
	}
)

func TestCreateUser(t *testing.T) {
	t.Run("when everything goes ok, should create a user register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		mock.ExpectBegin()
		mock.
			ExpectQuery(createUser).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(defaultID))
		mock.ExpectCommit()

		r := repository.New(dbGorm)
		result, err := r.User.CreateUser(ctx, user)

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
		mock.ExpectQuery(createUser).WillReturnError(wantErr)
		mock.ExpectRollback()

		r := repository.New(dbGorm)
		result, err := r.User.CreateUser(ctx, user)
		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, result)
	})
}

func TestPartialUpdate(t *testing.T) {
	t.Run("when everything goes ok, should update a user register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		mock.ExpectBegin()
		mock.
			ExpectExec(updatePartial).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		r := repository.New(dbGorm)
		err := r.User.PartialUpdateUser(ctx, partialUpdated)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.ExpectBegin()
		mock.
			ExpectExec(updatePartial).WillReturnError(wantErr)
		mock.ExpectRollback()

		r := repository.New(dbGorm)
		err := r.User.PartialUpdateUser(ctx, partialUpdated)

		assert.ErrorIs(t, err, wantErr)
	})
}
