package ui

import (
	h "github.com/mithcs/dof/internal/handlers"
	"github.com/urfave/cli/v3"
)

var deployCommand = &cli.Command{
	Name:      "deploy",
	Aliases:   []string{"d"},
	Usage:     "deploy files",
	ArgsUsage: "<names>",
	Arguments: []cli.Argument{
		&cli.StringArgs{
			Name: "names",
			Min:  1,
			Max:  -1,
		},
	},
	Action: h.DeployHandler,
}
