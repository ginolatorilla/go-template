package cmd

import (
	"fmt"

	"github.com/ginolatorilla/go-template/version"
	"github.com/spf13/cobra"
)

func (cli *CLI) buildVersionCommand() {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version:", version.Version)
			fmt.Println("Commit hash:", version.CommitHash)
		},
	}

	cli.rootCmd.AddCommand(command)
}
