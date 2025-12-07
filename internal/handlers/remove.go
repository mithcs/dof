package handlers

import (
	"context"

	"github.com/mithcs/dof/internal/files"
	md "github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// RemoveHandler is handler for remove subcommand
func RemoveHandler(ctx context.Context, cmd *cli.Command) error {
	m := md.Metadata{}
	name := cmd.String("name")
	entry, err := m.Get(name)
	if err != nil {
		return err
	}

	paths, err := files.AbsPaths(cmd.StringArgs("paths"))
	if err != nil {
		return err
	}

	paths = files.GeneralizePaths(paths)
	indices, err := pathIndices(paths, entry.Paths)
	if err != nil {
		return err
	}

	paths = files.ResolvePaths(paths)
	entry.Paths = removePaths(entry.Paths, indices)
	err = m.Update(entry)
	if err != nil {
		return err
	}

	err = files.MoveFromName(paths, indices, name)
	if err != nil {
		return err
	}

	return nil
}

// pathIndices returns indices of <toCheck> from <paths>.
// returns error if any of <toCheck> is not found in <paths>
func pathIndices(toCheck []string, paths []string) ([]int, error) {
	indices := []int{}

OUTER:
	for _, tc := range toCheck {
		for i, p := range paths {

			if tc == p {
				indices = append(indices, i)
				continue OUTER
			}
		}

		// TODO: also mention which path was not found
		return indices, md.ErrPathNotFound
	}

	return indices, nil
}

// removePaths removes paths from <paths> based on <indices>
func removePaths(paths []string, indices []int) []string {
	for _, idx := range indices {
		paths = append(paths[:idx], paths[idx+1:]...)
	}

	return paths
}
