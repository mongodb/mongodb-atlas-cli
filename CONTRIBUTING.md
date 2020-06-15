# Contributing to MongoDB CLI

Thanks for your interest in contributing to `mongocli`, 
this document describes some of the guidelines necessary to participate in the comunity 

## Feature Requests

We welcome any feedback or feature request, to submit yours
please head over to our [feedback page](https://feedback.mongodb.com/). 
 
## Reporting Issues

Please create a [GitHub issue](https://github.com/mongodb/mongocli/issues) describing the kind of problem you're facing
with as much detail as possible, including things like operating system or anything else that may be relevant to the issue.

## Submitting a Patch

Before submitting a patch to the repo please consider opening an [issue first](#reporting-issues)

### Contributor License Agreement

For patches to be accepted, contributors must sign our [CLA](https://www.mongodb.com/legal/contributor-agreement).

### Development Setup

#### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.14)](https://golang.org/dl/)

#### Environment
- Fork the repository.
- Clone your forked repository locally.
- We use Go Modules to manage dependencies, so you can develop outside of your `$GOPATH`.
- We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code, you can install it locally via `make setup`.

### Building and Testing

The following is a short list of commands that can be run in the root of the project directory

- Run `make` see a list of available targets.
- Run `make test` to run all unit tests.
- Run `make lint` to validate against our linting rules.
- Run `E2E_TAGS=e2e,atlas make e2e-test` will run end to end tests against an Atlas instance,
  please make sure to have set `MCLI_*` variables pointing to that instance.
- Run `E2E_TAGS=e2e,cloudmanager make e2e-test` will run end to end tests against an Cloud Manager instance.<br />
  Please remember to update `the e2e/cloud_manager/e2e.env` with the name of the host that is running your automation agent and 
  to set `MCLI_*` variables to point to your Cloud Manager instance. 
- Run `make build` to generate a local binary in the `./bin` folder.

We provide a git pre-commit hook to format and check the code, to install it run `make link-git-hooks` 

### Generating Mocks

We use [mockgen](https://github.com/golang/mock) to handle mocking in our unit tests.
If you need a new mock please update or add the `//go:generate` instruction to the appropriate file 

### Adding a New Command

`mongocli` uses [Cobra](https://github.com/spf13/cobra) as a framework for defining commands,
in addition to this we have defined a basic structure that should be followed.
For a `mongocli scope newCommand` command a file `internal/cli/scope_new_command.go` should implement: 
- A `ScopeNewCommandOpts` struct which handles the different options for the command.
- At least a `func (opts *ScopeNewCommandOpts) Run() error` function with the main command logic.
- A `func ScopeNewCommandBuilder() *cobra.Command` function to put together the expected cobra definition along with the `ScopeNewCommandOpts` logic.

### Third Party Dependencies

We scan our dependencies for vulnerabilities and incompatible licenses using [Snyk](https://snyk.io/).
To run Snyk locally please follow their [CLI reference](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI) 

## Maintainer's Guide

Reviewers, please ensure that the CLA has been signed by referring to [the contributors tool](https://contributors.corp.mongodb.com/) (internal link).
