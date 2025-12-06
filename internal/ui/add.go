package ui

import (
	"context"

	h "github.com/mithcs/dof/internal/handlers"
	"github.com/urfave/cli/v3"
)

var addCommand = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "add files to dof",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "set name for dotfiles",
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "copy",
			Aliases: []string{"c"},
		},
		&cli.BoolFlag{
			Name:    "symlink",
			Aliases: []string{"s"},
		},
	},
	ArgsUsage: "<dotfiles>",
	Arguments: []cli.Argument{
		&cli.StringArgs{
			Name: "paths",
			Min:  1,
			Max:  -1,
		},
	},
	Action:        h.AddHandler,
	ShellComplete: addShellCompletion,
}

func addShellCompletion(ctx context.Context, c *cli.Command) {
	// do not provide any completion for add subcommand
	// let the shell handle it
}
