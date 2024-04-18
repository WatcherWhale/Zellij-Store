package cli

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/zellij-store/pkg/shell"
)

func Run(ctx context.Context, session string) error {
	app := cli.App{
		Name: "zellij-store",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "session",
				Value: session,
				Usage: "The Zellij session, default value is current session.",
			},
			&cli.StringFlag{
				Name:  "shell",
				Value: shell.DetectShell(),
				Usage: "The shell syntax to set variables with.",
			},
		},
		Commands: []*cli.Command{
			&storeCommand,
			&loadCommand,
			&cleanCommand,
		},
		Action: loadCommandAction,
	}

	return app.RunContext(ctx, os.Args)
}
