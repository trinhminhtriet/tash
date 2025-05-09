package cmd

import (
	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/tash/core/dao"
)

func describeCmd(config *dao.Config, configErr *error) *cobra.Command {
	cmd := cobra.Command{
		Aliases: []string{"desc"},
		Use:     "describe <servers|tasks>",
		Short:   "Describe servers, tasks, specs and targets",
		Long:    "Describe servers, tasks, specs and targets",
		Example: `  # Describe servers
  tash describe servers

  # Describe tasks
  tash describe tasks`,
		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		describeServersCmd(config, configErr),
		describeTasksCmd(config, configErr),
		describeTargetsCmd(config, configErr),
		describeSpecsCmd(config, configErr),
	)

	return &cmd
}
