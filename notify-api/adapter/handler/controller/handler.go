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
	notifyRoute := routes.Group("/notification")
	notifyRoute.POST("/:user_id", h.CreateNotification)
	notifyRoute.GET("", h.FetchNotification)
	notifyRoute.GET("/message", h.FetchMessage)
	notifyRoute.POST("/message", h.SendMessage)
}
