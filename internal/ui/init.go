package ui

import (
	h "github.com/mithcs/dof/internal/handlers"
	"github.com/urfave/cli/v3"
)

var initCommand = &cli.Command{
	Name:   "init",
	Usage:  "initialize dof",
	Action: h.InitHandler,
}
