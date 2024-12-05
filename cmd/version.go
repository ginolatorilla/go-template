package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	AppName    = "go-template" // Name of the application
	Version    = ""            // Version of the application
	CommitHash = ""            // Commit hash of the application
)

// buildVersionCommand creates the version command.
func (cli *CLI) buildVersionCommand() {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version:   ", Version)
			fmt.Println("Commit:    ", CommitHash)
			fmt.Println("Go Version:", runtime.Version())
			fmt.Println("Compiler:  ", runtime.Compiler)
			fmt.Println("Platform:  ", runtime.GOOS, runtime.GOARCH)
		},
	}

	cli.rootCmd.AddCommand(command)
}
