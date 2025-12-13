package handlers

import (
	"context"
	"errors"

	"github.com/mithcs/dof/internal/config"
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

	fallback, err := methodFromConfig()
	if err != nil {
		return err
	}
	method := selectMethod(entry.Method, fallback, cmd)

	pathsCount := len(entry.Paths)
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

// selectMethod selects method to use. It returns given method if set, otherwise
// checks command line flags to determine. If no flags are set either, fallback
// method is returned
func selectMethod(method md.Method, fallback md.Method, cmd *cli.Command) md.Method {
	if method != "" {
		return method
	}

	switch {
	case cmd.Bool("copy"):
		return md.Copy
	case cmd.Bool("symlink"):
		return md.Symlink
	default:
		return fallback
	}
}

// addFile adds file based on method
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

// methodFromConfig returns method from config
func methodFromConfig() (md.Method, error) {
	c := &config.Config{}
	method, err := c.DefaultMethod()
	if err != nil {
		return md.Copy, err
	}

	if method == "symlink" {
		return md.Symlink, err
	}

	return md.Copy, err
}
