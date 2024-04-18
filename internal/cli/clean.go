package cli

import (
	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/zellij-store/pkg/store"
)

var cleanCommand cli.Command = cli.Command{
	Name:  "clean",
	Usage: "",
	Aliases: []string{
		"c",
	},
	Action: cleanCommandAction,
}

func cleanCommandAction(ctx *cli.Context) error {
	return store.CleanCache()
}
