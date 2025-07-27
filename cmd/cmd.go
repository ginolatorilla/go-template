// Package cmd provides the command line interface for the application.
//
// Copyright © 2024 Gino Latorilla
package cmd

import (
	"github.com/ginolatorilla/core"
	"github.com/ginolatorilla/go-template/cmd/root"
	"github.com/ginolatorilla/go-template/cmd/version"
)

var (
	AppName    = "go-template" // Name of the application
	Version    = ""            // Version of the application
	CommitHash = ""            // Commit hash of the application
)

func Execute() {
	cmd := root.NewCommand(AppName)
	cmd.AddCommand(version.NewCommand(Version, CommitHash))
	core.Check(cmd.Execute())
}
