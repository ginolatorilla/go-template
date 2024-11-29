/*
Package version contains the build metadata of the application.

These variables can be overriden with linker flags during the build process.
  - AppName
  - Version
  - CommitHash
*/
package version

var (
	AppName    = "go-template" // Name of the application
	Version    = ""            // Version of the application
	CommitHash = ""            // Commit hash of the application
)
