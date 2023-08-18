package controller

import (
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase port.NotifyUseCase
}

func NewHandler(n port.NotifyUseCase) *Handler {
	return &Handler{
		useCase: n,
	}
}

func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	userRoute := routes.Group("/notification")
	userRoute.POST("/:user_id", h.CreateNotification)
}
