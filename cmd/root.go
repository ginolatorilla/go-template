// Package cmd provides the command line interface for the application.
//
// Copyright © 2024 Gino Latorilla
package cmd

import (
	"fmt"
	"os"

	"github.com/ginolatorilla/go-template/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CLI represents the command line interface for the application.
type CLI struct {
	rootCmd    *cobra.Command
	configFile string
	verbosity  int
}

// NewCLIApp creates a new CLI application.
//
// This is where you need to call every other build*Command method.
func NewCLIApp() (*CLI, error) {
	var cli CLI
	cli.buildRootCommand()

	// Register commands here. The build*Command functions must register themselves with the root command.
	cli.buildVersionCommand()
	return &cli, nil
}

// Execute runs the CLI application.
//
// This parses all command line arguments and executes the command's Run function.
func (cli *CLI) Execute() error {
	return cli.rootCmd.Execute()
}

// buildRootCommand creates the root command for the CLI application.
func (cli *CLI) buildRootCommand() {
	cli.rootCmd = &cobra.Command{
		Use:   version.AppName,
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	cli.rootCmd.PersistentFlags().StringVar(
		&cli.configFile,
		"config",
		"",
		fmt.Sprintf("Read configuration from this file (default is $HOME/.%s.yaml)", version.AppName),
	)
	cli.rootCmd.PersistentFlags().CountVarP(
		&cli.verbosity,
		"verbose",
		"v",
		"Verbosity level. Use -v for verbose, -vv for more verbose, etc.",
	)

	// Have to be here because we need to run these functions when the root command parses its flags
	// and before it runs any subcommands.
	cobra.OnInitialize(cli.setUpLogger, cli.configure)
}

// setUpLogger sets up the logger based on the verbosity level.
//
// This function mimics the default logging level of Python's logger (starts at WARNING).
func (cli *CLI) setUpLogger() {
	var level zapcore.Level
	switch cli.verbosity {
	case 0:
		level = zap.WarnLevel
	case 1:
		level = zap.InfoLevel
	default:
		level = zap.DebugLevel
	}

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	logger := zap.Must(config.Build())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

// configure reads application options from a file and environment variables.
func (cli *CLI) configure() {
	if cli.configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cli.configFile)
	} else {
		zap.S().Debug("No config file specified, searching for default config file")
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			zap.S().Fatalf("failed to get user home directory: %w", err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", version.AppName))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		zap.S().Info("Using config file:", viper.ConfigFileUsed())
	}
}
