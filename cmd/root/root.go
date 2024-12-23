package root

import (
	"fmt"
	"os"

	u "github.com/ginolatorilla/go-template/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewCommand(appName string) *cobra.Command {
	var configFile string
	var verbosity int

	cobra.OnInitialize(
		func() { setUpLogger(verbosity) },
		func() { configure(configFile, appName) },
	)

	cmd := &cobra.Command{
		Use:   appName,
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// This function is a placeholder only to visually check if logging works.
			defer zap.S().Sync()
			zap.S().Warn("Hello, World!")
			zap.S().Info("Hello, World!")
			zap.S().Debug("Hello, World!")
		},
	}

	cmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		fmt.Sprintf("Read configuration from this file (default is $HOME/.%s.yaml)", appName),
	)
	cmd.PersistentFlags().CountVarP(
		&verbosity,
		"verbose",
		"v",
		"Verbosity level. Use -v for verbose, -vv for more verbose, etc.",
	)
	return cmd
}

// setUpLogger sets up the logger based on the verbosity level.
//
// This function mimics the default logging level of Python's logger (starts at WARNING).
func setUpLogger(verbosity int) {
	lvl := zap.WarnLevel
	trace := false

	switch verbosity {
	case 0:
		lvl = zap.WarnLevel
		trace = false
	case 1:
		lvl = zap.InfoLevel
		trace = false
	case 2:
		lvl = zap.DebugLevel
		trace = false
	default:
		lvl = zap.DebugLevel
		trace = true
	}

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.DisableStacktrace = !trace
	zap.ReplaceGlobals(zap.Must(config.Build()))
}

// configure reads application options from a file and environment variables.
func configure(configFile, appName string) {
	if configFile != "" {
		viper.SetConfigFile(configFile)
		return
	}

	zap.S().Debug("No config file specified, searching for default config file")
	home := u.Must(os.UserHomeDir())
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(fmt.Sprintf(".%s", appName))

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		zap.S().Info("Using config file:", viper.ConfigFileUsed())
	}
}
