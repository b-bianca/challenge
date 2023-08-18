package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/b-bianca/melichallenge/adapter/handler/user"
	"github.com/b-bianca/melichallenge/adapter/model"
	"github.com/b-bianca/melichallenge/internal/domain/entity"
	"github.com/b-bianca/melichallenge/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createModelInput = &model.UserRequest{
		Name:     "test",
		CPF:      "12312312312",
		Email:    "t@email.com",
		Password: "pass14**",
	}

	createEntityInput = &entity.User{
		Name:     "test",
		CPF:      "12312312312",
		Email:    "t@email.com",
		Password: "pass14**",
	}

	createEntityOutput = &entity.User{
		ID:           uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		Name:         "test",
		CPF:          "12312312312",
		Email:        "t@email.com",
		Notification: true,
	}

	createModelOutput = &model.UserResponse{
		ID:           uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		Name:         "test",
		CPF:          "12312312312",
		Email:        "t@email.com",
		Notification: true,
	}

	updatePath = "/user/8c2b51bf-7b4c-4a4b-a024-f283576cf190"

	updateModelInput = &model.OptoutRequest{
		Notification: false,
	}

	updateEntityInput = &entity.User{
		ID:           uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		Notification: false,
	}
)

func TestCreateUser(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
		fmt.Println(jsonBytes)
		if err != nil {
			return
		}
		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		ctxGin, _ := gin.CreateTestContext(w)
		ctxGin.Request = req

		usecaseMock := mocks.NewUserUseCase(t)
		usecaseMock.On("CreateUser", ctxGin, createEntityInput).Return(createEntityOutput, nil).Once()

		handler := user.NewHandler(usecaseMock)

		handler.CreateUser(ctxGin)

		res := w.Result()
		defer res.Body.Close()
		got, err := json.Marshal(createModelOutput)

		assert.NoError(t, err)
		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(got))
		assert.Equal(t, http.StatusOK, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{}`)))
		w := httptest.NewRecorder()
		ctxGin, _ := gin.CreateTestContext(w)
		ctxGin.Request = req

		usecaseMock := mocks.NewUserUseCase(t)

		handler := user.NewHandler(usecaseMock)

		handler.CreateUser(ctxGin)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when use case return error; should return response 500", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
		fmt.Println(jsonBytes)
		if err != nil {
			return
		}
		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()
		ctxGin, _ := gin.CreateTestContext(w)
		ctxGin.Request = req

		usecaseMock := mocks.NewUserUseCase(t)
		usecaseMock.On("CreateUser", ctxGin, createEntityInput).Return(nil, errors.New("error")).Once()

		handler := user.NewHandler(usecaseMock)

		handler.CreateUser(ctxGin)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})
}

func TestOptoutUser(t *testing.T) {
	t.Run("when everything is ok, should return response 204", func(t *testing.T) {
		jsonBytes, err := json.Marshal(updateModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPatch, updatePath, bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewUserUseCase(t)
		usecaseMock.On("OptoutUser", mock.AnythingOfType("*gin.Context"), updateEntityInput).Return(nil).Once()

		handler := user.NewHandler(usecaseMock)

		assert.NoError(t, err)
		engine.PATCH("/user/:id", handler.OptoutUser)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, updatePath, bytes.NewBuffer([]byte(`>`)))

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewUserUseCase(t)
		handler := user.NewHandler(usecaseMock)

		engine.PATCH("/user/:id", handler.OptoutUser)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("when use case return an error, should return response 500", func(t *testing.T) {
		jsonBytes, err := json.Marshal(updateModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPatch, updatePath, bytes.NewBuffer(jsonBytes))

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewUserUseCase(t)
		handler := user.NewHandler(usecaseMock)

		usecaseMock.On("OptoutUser", mock.AnythingOfType("*gin.Context"), updateEntityInput).Return(errors.New("error")).Once()

		engine.PATCH("/user/:id", handler.OptoutUser)
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
