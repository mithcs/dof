package ui

import (
	"context"
	"fmt"

	h "github.com/mithcs/dof/internal/handlers"
	md "github.com/mithcs/dof/internal/metadata"
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
	Action:        h.DeployHandler,
	ShellComplete: deployShellCompletion,
}

func deployShellCompletion(ctx context.Context, c *cli.Command) {
	m := &md.Metadata{}
	entries, err := m.All()
	if err != nil {
		return
	}

	for _, e := range entries {
		fmt.Println(e.Name)
	}
}
