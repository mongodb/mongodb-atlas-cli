# Contributing to the CLI

Thanks for your interest in contributing to `mcli`, 
this document describe the necessary steps to get a development environment going and the best way to contribute back to the project

## Development setup

### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.13)](https://golang.org/dl/)

### Environment
- Fork the repository.
- Clone your forked repository locally.
- We use Go Modules to manage dependencies, so you can develop outside of your `$GOPATH`.

We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code, you can install it locally via `make setup`.

## Building and testing

The following is a short list of commands that can be run in the root of the project directory

- Run `make` to generate a local binary in the `./bin` folder.
- Run `make test` to run all unit tests.
- Run `make lint` to validate against our linting rules.

We provide a git pre-commit hook to format and check the code, to install it run `make link-git-hooks` 

### Generating mocks

We use [mockgen](https://github.com/golang/mock) to handle mocking in our unit tests
If you need a new mock please add a reference on the [Make](Makefile) file and run `make gen-mocks`

### Adding a new command

`mcli` uses [Cobra](https://github.com/spf13/cobra) as a framework for defining commands,
in addition to this we have defined a basic structure that should be followed.
For a `mcli scope newCmd` command a file `internal/cli/scope_new_command.go` should implement: 
- A `ScopeNewCommandOpts` struct which handles the different options for the command.
- At least a `func (opts *ScopeNewCommandOpts) Run() error` function with the main command logic.
- A `func ScopeNewCommandBuilder() *cobra.Command` function to put together the expected cobra definition along with the `ScopeNewCommandOpts` logic.
