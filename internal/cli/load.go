package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/zellij-store/pkg/shell"
	"github.com/watcherwhale/zellij-store/pkg/store"
)

var loadCommand cli.Command = cli.Command{
	Name:  "load",
	Usage: "Load env variables",
	Aliases: []string{
		"l",
	},
	Action: loadCommandAction,
}

func loadCommandAction(ctx *cli.Context) error {
	store, err := store.LoadStore(ctx.String("session"))
	if err != nil {
		return nil
	}

	script := shell.GenerateScript(ctx.String("shell"), store)
	fmt.Println(script)

	return nil
}
