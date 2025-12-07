package ui

import (
	h "github.com/mithcs/dof/internal/handlers"
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
	Action: h.RemoveHandler,
}
