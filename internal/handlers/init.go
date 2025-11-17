package handlers

import (
	"context"

	"github.com/mithcs/dof/internal/config"
	"github.com/mithcs/dof/internal/files"
	"github.com/mithcs/dof/internal/metadata"
	"github.com/urfave/cli/v3"
)

// InitHandler is handler for init subcommand
func InitHandler(ctx context.Context, cmd *cli.Command) error {
	err := files.CreateDotfilesDir()
	if err != nil {
		return err
	}

	err = files.CreateDofDir()
	if err != nil {
		return err
	}

	m := &metadata.Metadata{}
	err = m.Create()
	if err != nil {
		return err
	}

	c := &config.Config{}
	err = c.Create()
	if err != nil {
		return err
	}

	return nil
}
