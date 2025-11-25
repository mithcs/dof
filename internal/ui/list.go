package ui

import (
	h "github.com/mithcs/dof/internal/handlers"
	"github.com/urfave/cli/v3"
)

var listCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list files managed by dof",
	Action:  h.ListHandler,
}
