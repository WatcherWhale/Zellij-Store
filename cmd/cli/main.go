package main

import (
	"context"
	"fmt"
	"os"

	"github.com/watcherwhale/zellij-store/internal/cli"
)

func main() {
	session := os.Getenv("ZELLIJ_SESSION_NAME")
	err := cli.Run(context.Background(), session)

	if err != nil {
		fmt.Printf("An error occurred while running: %v", err)
		os.Exit(1)
	}
}
