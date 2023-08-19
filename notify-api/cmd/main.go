package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	manage "github.com/b-bianca/melichallenge/notify-api/adapter/handler"
	"github.com/b-bianca/melichallenge/notify-api/adapter/repository"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const (
	shutdownTimeout = 5 * time.Second
	pathPrefix      = "/api/v1"
)

func main() {

	repository := repository.NewRepository()

	m := manage.New(&manage.UseCases{
		Notify: usecase.NewNotifyUseCase(repository.Notify),
	})

	engine := gin.Default()

	v1Routes := engine.Group(pathPrefix)

	m.RegisterRoutes(v1Routes)

	engine.Run(":8081")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: engine,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
	}
}
