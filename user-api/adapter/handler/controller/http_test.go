package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	user "github.com/b-bianca/melichallenge/user-api/adapter/handler/controller"
	"github.com/b-bianca/melichallenge/user-api/adapter/model"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createModelInput = &model.UserRequest{
		CPF: "12312312312",
	}

	createEntityInput = &entity.User{
		CPF: "12312312312",
	}

	createEntityOutput = &entity.User{
		ID:           uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		CPF:          "12312312312",
		Notification: true,
	}

	createModelOutput = &model.UserResponse{
		ID:           uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		CPF:          "12312312312",
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

func TestPartialUpdateUser(t *testing.T) {
	t.Run("when everything is ok, should return response 204", func(t *testing.T) {
		jsonBytes, err := json.Marshal(updateModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPatch, updatePath, bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewUserUseCase(t)
		usecaseMock.On("PartialUpdateUser", mock.AnythingOfType("*gin.Context"), updateEntityInput).Return(nil).Once()

		handler := user.NewHandler(usecaseMock)

		assert.NoError(t, err)
		engine.PATCH("/user/:id", handler.PartialUpdateUser)
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

		engine.PATCH("/user/:id", handler.PartialUpdateUser)
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

		usecaseMock.On("PartialUpdateUser", mock.AnythingOfType("*gin.Context"), updateEntityInput).Return(errors.New("error")).Once()

		engine.PATCH("/user/:id", handler.PartialUpdateUser)
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
