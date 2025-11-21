package handlers

import (
	"context"
	"errors"

	"github.com/mithcs/dof/internal/files"
	"github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// AddHandler is handler for add subcommand
func AddHandler(ctx context.Context, cmd *cli.Command) error {
	// TODO: Cleanup
	name := cmd.String("name")

	m := &metadata.Metadata{}
	entry, err := m.Get(name)
	if err != nil && !errors.Is(err, metadata.ErrEntryNotFound) {
		return err
	}

	paths, err := files.AbsPaths(cmd.StringArgs("paths"))
	if err != nil {
		return err
	}

	start := len(entry.Paths)
	var addFile func([]string, string, int) error
	switch {
	case cmd.Bool("copy"):
		addFile = files.CopyToName
		entry.Method = metadata.Copy
	case cmd.Bool("symlink"):
		addFile = files.MoveAndSymlink
		entry.Method = metadata.Symlink
	default:
		addFile = files.CopyToName
		entry.Method = metadata.Copy
	}

	err = addFile(paths, name, start)
	if err != nil {
		return err
	}

	paths = files.GeneralizePaths(paths)
	entry.Name = name
	entry.Paths = append(entry.Paths, paths...)

	// if entry (we got earlier) is empty then Add() otherwise Update()
	if start == 0 {
		if err = m.Add(entry); err != nil {
			return err
		}
	} else {
		if err = m.Update(entry); err != nil {
			return err
		}
	}

	return nil
}
