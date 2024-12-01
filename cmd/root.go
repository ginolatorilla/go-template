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

type CLI struct {
	rootCmd    *cobra.Command
	configFile string
	verbosity  int
}

func NewCLIApp() (*CLI, error) {
	var cli CLI
	cli.buildRootCommand()
	cli.buildVersionCommand()
	cli.setUpLogger()
	if err := cli.configure(); err != nil {
		return nil, fmt.Errorf("failed to configure CLI: %w", err)
	}
	return &cli, nil
}

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
}

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

func (cli *CLI) configure() error {
	if cli.configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cli.configFile)
	} else {
		zap.S().Debug("No config file specified, searching for default config file")
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home directory: %w", err)
		}

		// Search config in home directory with name ".go-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", version.AppName))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		zap.S().Info("Using config file:", viper.ConfigFileUsed())
	}

	return nil
}

func (cli *CLI) Execute() error {
	return cli.rootCmd.Execute()
}
