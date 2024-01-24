package main

import (
	"os"

	"github.com/Arif9878/oauth2/cmd/http"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "http",
				Usage:  "http server",
				Action: http.Exec,
			},
		},
		Name:  "OAUTH2",
		Usage: "OAUTH2 server",
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
