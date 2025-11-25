package ui

import "github.com/urfave/cli/v3"

var SubCommands = []*cli.Command{
	initCommand,
	addCommand,
	listCommand,
}
