/*
Copyright © 2024 Gino Latorilla

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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

// rootCmdOptions contains the parsed command line arguments, which can be used by commands.
var rootCmdOptions struct {
	ConfigFile string
	Verbosity  int
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   version.AppName,
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger, initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&rootCmdOptions.ConfigFile,
		"config",
		"",
		fmt.Sprintf("Read configuration from this file (default is $HOME/.%s.yaml)", version.AppName),
	)
	rootCmd.PersistentFlags().CountVarP(
		&rootCmdOptions.Verbosity,
		"verbose",
		"v",
		"Verbosity level. Use -v for verbose, -vv for more verbose, etc.",
	)
}

func initLogger() {
	var level zapcore.Level
	switch rootCmdOptions.Verbosity {
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if rootCmdOptions.ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(rootCmdOptions.ConfigFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", version.AppName))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
