// Package cmd provides the command line interface for the application.
//
// Copyright © 2024 Gino Latorilla
package cmd

import (
	"github.com/ginolatorilla/go-template/cmd/root"
	"github.com/ginolatorilla/go-template/cmd/version"

	u "github.com/ginolatorilla/go-template/pkg/utils"
)

var (
	AppName    = "go-template" // Name of the application
	Version    = ""            // Version of the application
	CommitHash = ""            // Commit hash of the application
)

func Execute() {
	cmd := root.NewCommand(AppName)
	cmd.AddCommand(version.NewCommand(Version, CommitHash))
	u.Check(cmd.Execute())
}
