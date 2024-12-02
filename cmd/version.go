package cmd

import (
	"fmt"
	"runtime"

	"github.com/ginolatorilla/go-template/version"
	"github.com/spf13/cobra"
)

// buildVersionCommand creates the version command.
func (cli *CLI) buildVersionCommand() {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version:   ", version.Version)
			fmt.Println("Commit:    ", version.CommitHash)
			fmt.Println("Go Version:", runtime.Version())
			fmt.Println("Compiler:  ", runtime.Compiler)
			fmt.Println("Platform:  ", runtime.GOOS, runtime.GOARCH)
		},
	}

	cli.rootCmd.AddCommand(command)
}
