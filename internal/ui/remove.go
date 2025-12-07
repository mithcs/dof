package ui

import (
	"context"
	"fmt"

	"github.com/mithcs/dof/internal/files"
	h "github.com/mithcs/dof/internal/handlers"
	md "github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

var removeCommand = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "remove files from dof",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "set name for dotfiles",
			Required: true,
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
	Action:        h.RemoveHandler,
	ShellComplete: removeShellCompletion,
}

func removeShellCompletion(ctx context.Context, c *cli.Command) {
	m := &md.Metadata{}
	entries, err := m.All()
	if err != nil {
		return
	}

	for _, e := range entries {
		e.Paths = files.ResolvePaths(e.Paths)

		for _, p := range e.Paths {
			fmt.Printf("%s:%s\n", p, e.Name)
		}
	}
}
