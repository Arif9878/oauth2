package http

import (
	"github.com/Arif9878/common/go/http"
	echoserver "github.com/Arif9878/common/go/http/echo/server"
	"github.com/Arif9878/oauth2/cmd/server"
	"github.com/Arif9878/oauth2/config"
	"github.com/Arif9878/oauth2/endpoints"
	"github.com/Arif9878/oauth2/endpoints/middleware"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func Exec(cCtx *cli.Context) error {
	ctx := cCtx.Context
	zerolog.Ctx(ctx)
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				config.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				validator.New,
			),
		),
		fx.Invoke(server.RunServers),
		fx.Invoke(middleware.ConfigMiddlewares),
		fx.Invoke(endpoints.ConfigEndpoints),
	).Run()
	zlog.Info().Msg("Running HTTP")

	return nil
}
