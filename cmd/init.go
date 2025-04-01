package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/tash/core"
	"github.com/trinhminhtriet/tash/core/dao"
	"github.com/trinhminhtriet/tash/core/print"
)

func initCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize tash in the current directory",
		Long:  "Initialize tash in the current directory.",
		Example: `  # Basic example
  tash init`,

		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			servers, err := dao.InitTash(args)
			core.CheckIfError(err)
			PrintServerInit(servers)
		},
		DisableAutoGenTag: true,
	}

	return &cmd
}

func PrintServerInit(servers []dao.Server) {
	theme := dao.Theme{
		Table: dao.DefaultTable,
	}

	options := print.PrintTableOptions{
		Theme:            theme,
		OmitEmptyRows:    true,
		OmitEmptyColumns: false,
		Output:           "table",
	}

	data := dao.TableOutput{
		Headers: []string{"server", "host"},
		Rows:    []dao.Row{},
	}

	for _, server := range servers {
		data.Rows = append(data.Rows, dao.Row{Columns: []string{server.Name, server.Host}})
	}

	fmt.Println("\nFollowing servers were added to tash.yaml")
	err := print.PrintTable(data.Rows, options, data.Headers, []string{}, true, true)
	core.CheckIfError(err)
}
