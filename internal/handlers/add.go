package handlers

import (
	"context"
	"errors"

	"github.com/mithcs/dof/internal/files"
	md "github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// AddHandler is handler for add subcommand
func AddHandler(ctx context.Context, cmd *cli.Command) error {
	m := &md.Metadata{}
	name := cmd.String("name")
	entry, err := m.Get(name)
	if err != nil && !errors.Is(err, md.ErrEntryNotFound) {
		return err
	}

	paths, err := files.AbsPaths(cmd.StringArgs("paths"))
	if err != nil {
		return err
	}

	pathsCount := len(entry.Paths)
	method := selectMethod(entry.Method, cmd)
	err = addFile(method, paths, name, pathsCount)
	if err != nil {
		return err
	}

	paths = files.GeneralizePaths(paths)

	entry.Name = name
	entry.Paths = append(entry.Paths, paths...)
	entry.Method = method

	err = saveEntry(m, pathsCount, entry)
	if err != nil {
		return err
	}

	return nil
}

// selectMethod selects method to use. It uses given method if set, otherwise
// checks flags to determine. If no flags are set either, default method is returned
func selectMethod(method md.Method, cmd *cli.Command) md.Method {
	if method != "" {
		return method
	}

	switch {
	case cmd.Bool("copy"):
		return md.Copy
	case cmd.Bool("symlink"):
		return md.Symlink
	default:
		// select default from config
		return md.Copy
	}
}

func addFile(method md.Method, paths []string, name string, pathCount int) error {
	if method == md.Symlink {
		return files.MoveAndSymlink(paths, name, pathCount)
	} else {
		return files.CopyToName(paths, name, pathCount)
	}
}

// saveEntry calls Add() method if pathsCount is 0, otherwise calls Update()
func saveEntry(m *md.Metadata, pathsCount int, e md.Entry) error {
	if pathsCount == 0 {
		return m.Add(e)
	} else {
		return m.Update(e)
	}
}
