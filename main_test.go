package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/ginolatorilla/go-template/version"
	"github.com/stretchr/testify/assert"
)

// Test_main runs the main function.
//
// The runMainInSubProcess function delegates to this function, passing arguments as necessary.
func Test_main(t *testing.T) {
	if os.Getenv("GO_TEST_SUBPROCESS") != "" {
		os.Args = append([]string{version.AppName}, extractMainSubProcessArgs(os.Args)...)
	}

	main()
}

func Test_main_ShouldAbortOnExecuteFailure(t *testing.T) {
	stdout, err := runMainInSubProcess("!!!unknown-command!!!")
	fmt.Println("---Begin output of main()---")
	fmt.Println(string(stdout))
	fmt.Println("---End output of main()---")
	assert.Error(t, err)
}

// runMainInSubProcess runs the main function in a subprocess.
func runMainInSubProcess(args ...string) ([]byte, error) {
	mainProc := exec.Command(os.Args[0], append([]string{"-test.run", "^Test_main$", "--"}, args...)...)
	mainProc.Env = append(os.Environ(), "GO_TEST_SUBPROCESS=1")
	return mainProc.CombinedOutput()
}

// extractMainSubProcessArgs extracts the arguments to pass to the main function in a subprocess.
//
// The input arguments, usually intended for go test, will be searched for the -- separator,
// and the arguments after it will be returned.
func extractMainSubProcessArgs(args []string) []string {
	for i, arg := range args {
		if arg == "--" {
			return args[i+1:]
		}
	}
	return nil
}
