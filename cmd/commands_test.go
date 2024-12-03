package cmd_test

import (
	"os"
	"testing"

	"github.com/ginolatorilla/go-template/cmd"
	"github.com/ginolatorilla/go-template/version"
	"github.com/stretchr/testify/assert"
)

func Test_VersionCommand(t *testing.T) {
	err := runCommand("version")
	assert.NoError(t, err)
}

func Test_Verbosity(t *testing.T) {
	for _, v := range []string{"-v", "-vv", "-vvv"} {
		t.Run(v, func(t *testing.T) {
			err := runCommand(v)
			assert.NoError(t, err)
		})
	}
}

func Test_Config(t *testing.T) {
	err := runCommand("--config", "testdata/config.yaml")
	assert.NoError(t, err)
}

func runCommand(args ...string) error {
	os.Args = append([]string{version.AppName}, args...)
	cli := cmd.NewCLIApp()
	return cli.Execute()
}
