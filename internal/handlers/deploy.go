package handlers

import (
	"context"

	"github.com/mithcs/dof/internal/files"
	md "github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// DeployHandler is handler for deploy subcommand
func DeployHandler(ctx context.Context, cmd *cli.Command) error {
	m := &md.Metadata{}
	names := cmd.StringArgs("names")

	entries, err := entriesFromNames(m, names)
	if err != nil {
		return err
	}

	err = deployFromEntries(entries)
	if err != nil {
		return err
	}

	return nil
}

// entriesFromNames returns metadata entry for each given name
func entriesFromNames(m *md.Metadata, names []string) ([]md.Entry, error) {
	entries := []md.Entry{}

	for _, n := range names {
		entry, err := m.Get(n)
		if err != nil {
			return entries, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// deployFromEntries deploys files based on metadata entries
func deployFromEntries(entries []md.Entry) error {
	var err error

	for _, e := range entries {
		e.Paths = files.ResolvePaths(e.Paths)

		if e.Method == md.Symlink {
			err = files.DeploySymlink(e.Paths, e.Name)
		} else {
			err = files.DeployCopy(e.Paths, e.Name)
		}
	}

	return err
}
