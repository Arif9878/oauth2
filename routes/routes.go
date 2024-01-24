package routes

import (
	"context"
	"net/http"

	"github.com/Arif9878/common/go/logger"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/health")
	group.GET("", healthCheck)
}

// Health Check api
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service is healthy!")
}
