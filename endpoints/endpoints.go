package endpoints

import (
	"context"

	"github.com/Arif9878/common/go/logger"
	"github.com/Arif9878/oauth2/routes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ConfigEndpoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {

	routes.MapRoute(validator, log, echo, ctx)
}
