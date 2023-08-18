package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/adapter/handler/controller"
	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createModelInput = &model.NotificationRequest{
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "test-message",
	}

	createEntityInput = &entity.Notification{
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "test-message",
	}

	createEntityOutput = &entity.Notification{
		ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf191"),
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "test-message",
	}

	postPath = "/notification/8c2b51bf-7b4c-4a4b-a024-f283576cf191"
)

func TestCreateNotification(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
		fmt.Println(jsonBytes)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		usecaseMock.On("CreateNotify", mock.AnythingOfType("*gin.Context"), createEntityInput).Return(createEntityOutput, nil).Once()

		handler := controller.NewHandler(usecaseMock)

		engine.POST("/notification/:user_id", handler.CreateNotification)
		engine.ServeHTTP(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer([]byte(`{>}`)))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		engine.POST("/notification/:user_id", handler.CreateNotification)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("when use case return an error, should return response 500", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer(jsonBytes))

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		usecaseMock.On("CreateNotify", mock.AnythingOfType("*gin.Context"), createEntityInput).Return(nil, errors.New("error")).Once()

		engine.POST("/notification/:user_id", handler.CreateNotification)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
