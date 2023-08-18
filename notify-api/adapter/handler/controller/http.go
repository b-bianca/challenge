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
