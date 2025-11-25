package handlers

import (
	"context"
	"fmt"

	"github.com/mithcs/dof/internal/files"
	md "github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// ListHandler is handler for list subcommand
func ListHandler(ctx context.Context, cmd *cli.Command) error {
	m := &md.Metadata{}
	entries, err := m.List()
	if err != nil {
		return err
	}

	printEntries(entries)

	return nil
}

// printEntries prints the entries
func printEntries(entries []md.Entry) {
	for _, e := range entries {
		fmt.Printf("%s (%s):\n", e.Name, e.Method)

		e.Paths = files.ResolvePaths(e.Paths)
		for _, p := range e.Paths {
			fmt.Println(p)
		}

		fmt.Println()
	}
}
