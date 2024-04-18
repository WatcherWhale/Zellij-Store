package cli

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/zellij-store/pkg/store"
)

var storeCommand cli.Command = cli.Command{
	Name:  "store",
	Usage: "Store an Environment Variable ",
	Aliases: []string{
		"s",
	},
	Args:      true,
	ArgsUsage: "A list of env variables to store",
	Action:    storeCommandAction,
}

func storeCommandAction(ctx *cli.Context) error {
	envSaves := ctx.Args().Slice()

	storeMap := map[string]string{}

	for _, env := range envSaves {
		storeMap[env] = os.Getenv(env)
	}

	err := store.SaveStore(ctx.String("session"), storeMap)
	if err != nil {
		return err
	}

	return nil
}
