# Gino's Go project template

A Go project template for command line applications.

## Features

- Command line parser with [Cobra](https://cobra.dev)
- Configuration management with [Viper](https://github.com/spf13/viper)
- Makefile with useful targets (see them by calling `make help`)
- Embed build information, such as version and git commit hashes, at build time
- Structured logging with [Zap](https://pkg.go.dev/go.uber.org/zap)
- Serve documentation generated from comments

## Requirements

- [Go](https://go.dev/doc/install)

## Demo

```shell
$ make
🌡  Running tests...
?       github.com/ginolatorilla/go-template    [no test files]
?       github.com/ginolatorilla/go-template/cmd        [no test files]
=== RUN   TestVersionsNotEmpty
--- PASS: TestVersionsNotEmpty (0.00s)
PASS
ok      github.com/ginolatorilla/go-template/version    (cached)
🧹 Tidying up package dependencies...
🏗️  Building the application...

$ go run main.go help -vv
2024-11-29T13:07:38.275+0800    DEBUG   cmd/root.go:107 No config file specified, searching for default config file
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  go-template [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version of the application

Flags:
      --config string   Read configuration from this file (default is $HOME/.go-template.yaml)
  -h, --help            help for go-template
  -v, --verbose count   Verbosity level. Use -v for verbose, -vv for more verbose, etc.

Use "go-template [command] --help" for more information about a command.
```

## Using this template

1. Copy all the files in this repo to your project.
2. Change the variables in the `Makefile` as indicated there by a TODO note.
3. Change the `LICENSE` file.
4. Change the default value of `AppName` in `version/version.go`.
5. Replace my name in the top-level comments in every `*.go` file.
6. Replace this `README.md` file.
