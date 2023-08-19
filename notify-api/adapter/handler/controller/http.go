package controller

import (
	"net/http"

	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateNotification(ctx *gin.Context) {
	userIDParam := ctx.Param("user_id")

	var input model.NotificationRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.Notification{
		UserID:   uuid.MustParse(userIDParam),
		DateTime: input.DateTime,
		Message:  input.Message,
	}

	res, err := h.useCase.CreateNotify(ctx, domain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	output := &model.NotificationResponse{
		ID:        res.ID,
		UserID:    uuid.MustParse(userIDParam),
		DateTime:  res.DateTime,
		Message:   res.Message,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) FetchNotification(ctx *gin.Context) {
	res, err := h.useCase.FetchNotify(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	responseItems := make([]*model.NotificationResponse, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &model.NotificationResponse{
			ID:        item.ID,
			UserID:    item.UserID,
			DateTime:  item.DateTime,
			Message:   item.Message,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	output := &model.NotificationListResponse{
		Result: responseItems,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) SendMessage(ctx *gin.Context) {
	var input model.MessageRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.Message{
		NotifyID: input.NotifyID,
		Message:  input.Message,
	}

	res, err := h.useCase.SendMessage(ctx, domain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	output := &model.MessageResponse{
		ID:        res.ID,
		NotifyID:  res.NotifyID,
		Message:   res.Message,
		CreatedAt: res.CreatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) FetchMessage(ctx *gin.Context) {
	res, err := h.useCase.FetchMessage(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	responseItems := make([]*model.MessageResponse, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &model.MessageResponse{
			ID:        item.ID,
			NotifyID:  item.NotifyID,
			Message:   item.Message,
			CreatedAt: item.CreatedAt,
		})
	}

	output := &model.MessageListResponse{
		Result: responseItems,
	}

	ctx.JSON(http.StatusOK, output)
}
