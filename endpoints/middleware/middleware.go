package middleware

import (
	"strings"

	"github.com/Arif9878/oauth2/models/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigMiddlewares(e *echo.Echo) {

	e.HideBanner = false

	e.Use(middleware.Logger())
	// e.HTTPErrorHandler = ProblemDetailsHandler
	// e.Use(otelmiddleware.EchoTracerMiddleware(jaegerCfg.ServiceName))

	// e.Use(echomiddleware.CorrelationIdMiddleware)
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constants.GzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

	e.Use(middleware.BodyLimit(constants.BodyLimit))
}
