package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	handler "github.com/b-bianca/melichallenge/user-api/adapter/handler/controller"
)

func TestRegisterRoutes(t *testing.T) {
	h := handler.NewHandler(nil)
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	h.RegisterRoutes(engine.Group("/api/v1"))

	routesInfo := engine.Routes()
	routesMethodAndPath := make([][]string, 0, len(routesInfo))
	for _, routeInfo := range routesInfo {
		routesMethodAndPath = append(routesMethodAndPath, []string{routeInfo.Method, routeInfo.Path})
	}

	expectedRoutesMethodAndPath := [][]string{
		{"PATCH", "/api/v1/user/:id"},
		{"POST", "/api/v1/user/"},
	}

	assert.Equal(t, expectedRoutesMethodAndPath, routesMethodAndPath)
}
