package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
		UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf191"),
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "test-message",
	}

	createEntityOutput = &entity.Notification{
		ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf191"),
		DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
		Message:  "test-message",
	}

	fetchEntityOutput = &entity.NotificationList{
		Result: []*entity.Notification{
			{
				ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
				UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf191"),
				DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
				Message:  "test-message",
			},
			{
				ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
				UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf193"),
				DateTime: time.Date(2023, 9, 04, 11, 00, 00, 00, time.UTC),
				Message:  "test-message-test",
			},
		},
	}

	fetchModelOutput = &model.NotificationListResponse{
		Result: []*model.NotificationResponse{
			{
				ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
				UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf191"),
				DateTime: time.Date(2023, 9, 03, 10, 00, 00, 00, time.UTC),
				Message:  "test-message",
			},
			{
				ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
				UserID:   uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf193"),
				DateTime: time.Date(2023, 9, 04, 11, 00, 00, 00, time.UTC),
				Message:  "test-message-test",
			},
		},
	}

	sendEntityInput = &entity.Message{
		NotifyID: uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf193"),
		Message:  "message",
	}

	sendModelInput = &model.MessageRequest{
		NotifyID: uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf193"),
		Message:  "message",
	}

	sendEntityOutput = &entity.Message{
		ID:       uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
		NotifyID: uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf193"),
		Message:  "message",
	}

	postPath        = "/notification/8c2b51bf-7b4c-4a4b-a024-f283576cf191"
	postMessagePath = "/notification/message"
	fetchPath       = "/notification"
)

func TestCreateNotification(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
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

func TestFetcNotification(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, fetchPath, nil)
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		usecaseMock.
			On("FetchNotify", mock.AnythingOfType("*gin.Context")).Return(fetchEntityOutput, nil).Once()

		handler := controller.NewHandler(usecaseMock)

		engine.GET("/notification", handler.FetchNotification)
		engine.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		wantGot, err := json.Marshal(fetchModelOutput)
		assert.NoError(t, err)

		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(wantGot))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("when use case return error; should return response 500", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, fetchPath, nil)
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)

		usecaseMock.
			On("FetchNotify", mock.AnythingOfType("*gin.Context")).Return(nil, errors.New("error")).Once()

		handler := controller.NewHandler(usecaseMock)

		engine.GET("/notification", handler.FetchNotification)
		engine.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		jsonBytes, err := json.Marshal(sendModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPost, postMessagePath, bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		usecaseMock.On("SendMessage", mock.AnythingOfType("*gin.Context"), sendEntityInput).Return(sendEntityOutput, nil).Once()

		handler := controller.NewHandler(usecaseMock)

		engine.POST("/notification/message", handler.SendMessage)
		engine.ServeHTTP(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, postMessagePath, bytes.NewBuffer([]byte(`{>}`)))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		engine.POST("/notification/message", handler.SendMessage)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("when use case return an error, should return response 500", func(t *testing.T) {
		jsonBytes, err := json.Marshal(sendModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPost, postMessagePath, bytes.NewBuffer(jsonBytes))

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewNotifyUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		usecaseMock.On("SendMessage", mock.AnythingOfType("*gin.Context"), sendEntityInput).Return(nil, errors.New("error")).Once()

		engine.POST("/notification/message", handler.SendMessage)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
