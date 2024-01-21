package server

import (
	"context"
	"errors"
	"net/http"

	echoserver "github.com/Arif9878/common/go/http/echo/server"
	"github.com/Arif9878/common/go/logger"
	"github.com/Arif9878/oauth2/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, logger logger.ILogger, e *echo.Echo, ctx context.Context, cfg *config.Config) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := echoserver.RunHttpServer(ctx, e, logger, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
					logger.Fatalf("error running http server: %v", err)
				}
			}()

			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, config.GetMicroserviceName(cfg.ServiceName))
			})

			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Infof("all servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}
