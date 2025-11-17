package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mithcs/dof/internal/ui"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:     "dof",
		Version:  "v0.0.2",
		Usage:    "Manage dot files easily",
		Commands: ui.SubCommands,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}
