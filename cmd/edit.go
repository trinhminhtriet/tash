package cmd

import (
	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/tash/core"
	"github.com/trinhminhtriet/tash/core/dao"
)

func editCmd(config *dao.Config, configErr *error) *cobra.Command {
	cmd := cobra.Command{
		Aliases: []string{"e", "ed"},
		Use:     "edit [flags]",
		Short:   "Open up tash config file in $EDITOR",
		Long:    "Open up tash config file in $EDITOR.",
		Example: `  # Edit current context
  tash edit`,
		Run: func(cmd *cobra.Command, args []string) {
			err := *configErr
			switch e := err.(type) {
			case *core.ConfigNotFound:
				core.CheckIfError(e)
			default:
				runEdit(*config)
			}
		},
		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		editServer(config, configErr),
		editTask(config, configErr),
		editTarget(config, configErr),
		editSpec(config, configErr),
	)

	return &cmd
}

func runEdit(config dao.Config) {
	err := config.EditConfig()
	core.CheckIfError(err)
}
