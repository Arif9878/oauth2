package main

import (
	"os"

	"github.com/Arif9878/oauth2/cmd/grpc"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "grpc",
				Usage:  "gRPC server",
				Action: grpc.Exec,
			},
		},
		Name:  "OAUTH2",
		Usage: "OAUTH2 server",
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
