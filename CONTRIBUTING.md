# Contributing to MongoDB CLI

Thanks for your interest in contributing to MongoDB Atlas CLI and MongoDB CLI,
this document describes some guidelines necessary to participate in the community.

## Table of Contents

- [Asking Support Questions](#asking-support-questions)
- [Feature Requests](#feature-requests)
- [Reporting Issues](#reporting-issues)
- [Auto-close stale issues and PRs](#auto-close-stale-issues-and-pull-requests)
- [Submitting Patches](#submitting-patches)
  - [Code Contribution Guidelines](#code-contribution-guidelines)
  - [Development Setup](#development-setup)
  - [Building and Testing](#building-and-testing)
  - [Adding a New Command](#adding-a-new-command)
  - [Third Party Dependencies](#third-party-dependencies)
- [Maintainer's Guide](#maintainers-guide)

## Asking Support Questions

MongoDB support is provided under MongoDB Atlas or Enterprise Advanced [support plans](https://support.mongodb.com/welcome).
Please don't use the GitHub issue tracker to ask questions.

## Feature Requests

We welcome any feedback or feature request, to submit yours
please head over to our [feedback page](https://feedback.mongodb.com/forums/930808-mongodb-cli).

## Reporting Issues

Please create a [GitHub issue](https://github.com/mongodb/mongodb-atlas-cli/issues/new?assignees=&labels=&template=bug_report.md) describing the kind of problem you're facing
with as much detail as possible, including things like operating system or anything else may be relevant to the issue.

## Auto-close Stale Issues and Pull Requests

- After 30 days of no activity (no comments or commits on an issue/PR) we automatically tag it as "stale" and add a message: ```This issue/PR has gone 30 days without any activity and meets the project's definition of "stale". This will be auto-closed if there is no new activity over the next 60 days. If the issue is still relevant and active, you can simply comment with a "bump" to keep it open, or add the label "not_stale". Thanks for keeping our repository healthy!```
- After 60 more days of no activity we automatically close the issue/PR.

## Submitting Patches

The Atlas CLI project welcomes all contributors and contributions regardless of skill or experience level.
If you are interested in helping with the project, please follow our [guidelines](#code-contribution-guidelines).

### Code Contribution Guidelines

To create the best possible product for our users and the best contribution experience for our developers,
we have a set of guidelines to ensure that all contributions are acceptable.

To make the contribution process as seamless as possible, we ask for the following:

- Fork the repository to work on your changes. Note that code contributions are accepted through pull requests to encourage discussion and allow for a smooth review experience.
- When you’re ready to create a pull request, be sure to:
  - Sign the [CLA](https://www.mongodb.com/legal/contributor-agreement).
  - Have test cases for the new code. If you have questions about how to do this, please ask in your pull request or check the [Building and Testing](#building-and-testing) section.
  - Run `make fmt`.
  - Add documentation if you are adding new features or changing functionality.
  - Confirm that `make check` succeeds. [GitHub Actions](https://github.com/mongodb/mongodb-atlas-cli/actions).

### Development Setup

#### Prerequisite Tools

- [Git](https://git-scm.com/)
- [Go (at least Go 1.20)](https://golang.org/dl/)

#### Environment

- Fork the repository.
- Clone your forked repository locally.
- We use Go Modules to manage dependencies, so you can develop outside your `$GOPATH`.
- We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code, you can install it locally via `make setup`.

### Building and Testing

The following is a short list of commands that can be run in the root of the project directory

- Run `make` see a list of available targets.
- Run `make test` to run all unit tests.
- Run `make lint` to validate against our linting rules.
- Run `E2E_TAGS=e2e,atlas make e2e-test` will run end to end tests against an Atlas instance,
  please make sure to have set `MCLI_*` variables pointing to that instance.
- Run `E2E_TAGS=cloudmanager,remote,generic make e2e-test` will run end-to-end tests against a Cloud Manager instance.<br />
  Please remember to: (a) have a running automation agent, and (b) set MCLI\_\* variables to point to your Cloud Manager instance.
- Run `make build` to generate a local binary in the `./bin` folder.

We provide a git pre-commit hook to format and check the code, to install it run `make link-git-hooks`.

#### Generating Mocks

We use [mockgen](https://github.com/golang/mock) to handle mocking in our unit tests.
If you need a new mock please update or add the `//go:generate` instruction to the appropriate file.

#### Compilation in VSCode

Please add following line to your settings.json file :
```
    "go.buildTags": "unit,e2e",
    "go.testTags": "unit,e2e"
```

This will enable compilation for unit test and end to end tests.

#### Debugging in VSCode

To debut in VSCode you need to create an debug configuration for the command with required arguments.
Run following commands to 

```
touch .vscode/launch.json
```
Then put following configuration into the file.
Review and replace command name and arguments depending on the command you are using.

```json
{
    "configurations": [
        {
            "name": "Login Command",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/atlas",
            "env": {},
            "buildFlags": "-ldflags '-X github.com/mongodb/mongodb-atlas-cli/internal/config.ToolName=atlascli'",
            "args": [
              "login"
            ]
      },
    ]
} 

```


### API Interactions

Atlas CLI and MongoDB CLI use [go-client-mongodb-atlas](https://github.com/mongodb/go-client-mongodb-atlas/) 
and [go-client-mongodb-ops-manager](https://github.com/mongodb/go-client-mongodb-ops-manager/) to interact with Atlas or Ops Manager/Cloud Manager.
Any new feature should first update the respective client.

### Adding a New Command

`atlascli` and `mongocli` use [Cobra](https://github.com/spf13/cobra) as a framework for defining commands,
in addition to this we have defined a basic structure that should be followed.
For a `mongocli scope newCommand` command, a file `internal/cli/scope/new_command.go` should implement:

- A `ScopeNewCommandOpts` struct which handles the different options for the command.
- At least a `func (opts *ScopeNewCommandOpts) Run() error` function with the main command logic.
- A `func ScopeNewCommandBuilder() *cobra.Command` function to put together the expected cobra definition along with the `ScopeNewCommandOpts` logic.
- A set of documentation fields further described in the section below.

Commands follow a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) approach to match the APIs, whenever possible.
For that reason, command arguments tend to match the path and query params of the APIs,
with the last param being a required argument and the rest handled via flag options.
For commands that create or modify complex data structures, the use of configuration files is preferred over flag options.

Note: we are experimenting with a generator, make sure to try it out in [tools/cli-generator](./tools/cli-generator/)

#### How to define flags:

Flags are a way to modify the command, also may be called "options". Flags always have a long version with two dashes (--state) but may also have a shortcut with one dash and one letter (-s).

`atlascli` uses the following types of flags:

- `--flagName value`: this type of flag passes the value to the command. Examples: `--projectId 5efda6aea3f2ed2e7dd6ce05`
- `--booleanFlag`: this flag represents a boolean and it sets the related variable to true when the flag is used, false otherwise. Example: `--force`
- `--flagName value1,value2,..,valueN`: you will also find flags that accept a list of values. This type of flag can be very useful to represent data structures as `--role roleName1@db,roleName2@db`, `--privilege action@dbName.collection,action2@dbName.collection`, or `--key field:type`.
  As shown in the examples, the standard format used to represent data structures consists of splitting the first value with the second one by at sign `@` or colon `:`, and the second value with the third one by a full stop `.`.
  We recommend using configuration files for complex data structures that require more than three values. For an example of configuration files, see [mongocli atlas cluster create](https://github.com/mongodb/mongodb-atlas-cli/blob/f2e6d661a3eb2cfcf9baab5f9e0b1c0f872b8c14/internal/cli/atlas/clusters/create.go#L235).

#### Documentation Requirements

If you are adding a brand-new command, or updating a command that has no doc annotations, please define the following doc structures for the command. For more information on all command structs, see [Cobra](https://pkg.go.dev/github.com/spf13/cobra#Command).

- Add `Use` - (Required) Shows the command and arguments if applicable. Will show up in 'help' output.
- Add `Short` - (Required) Briefly describes the command. Will show up in 'help' output.
- Add `Example` - (Required) Example of how to use the command. Will show up in 'help' output.
- Add `Annotations` - If the command has arguments, annotations should be added. They consist of key/value pairs that describe arguments in the command and are added to the generated documentation.
- Add `Long` - Fully describes the command. Will show up in 'help' output.

Furthermore, after adding the necessary structure, ensure that applicable documentation is generated by running `make gen-docs`.

- Run `make gen-docs`- This generates the documentation for the introduced command.
- Review the PR with the doc team.

### Third Party Dependencies

We scan our dependencies for vulnerabilities and incompatible licenses using [Snyk](https://snyk.io/).
To run Snyk locally please follow their [CLI reference](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI).

## Maintainer's Guide

Reviewers, please ensure that the CLA has been signed by referring to [the contributors tool](https://contributors.corp.mongodb.com/) (internal link).

For changes that involve user facing copy please include `docs-cloud-team` as a reviewer.

## SDK integration

Atlas CLI uses [atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) for API integration.
Go SDK will be automatically updated for the new versions using dependabot.
In situations when SDK does new major releases developers need to specify the version explicitly in the go update command. For example:

```
go get go.mongodb.org/atlas-sdk/v20230501001
```

Atlas CLI can work with multiple versions of the GO SDK supporting various Resource Versions. 

For more info please refer to the [SDK documentation](https://github.com/mongodb/atlas-sdk-go/blob/main/docs/doc_1_concepts.md#release-strategy-semantic-versioning) and 
[golang documentation](https://go.dev/doc/modules/version-numbers#major)

### Major Version Updates   

When adding a new major version of the go sdk, the old sdk version dependency will be still present in the go mod files.
Atlas CLI developers should update all imports to new major versions and remove old dependencies.

To update simply rename all instances of major version across the repository imports and go.mod files.

e.g `v20230201001` => `v20230201002` 

### Stable Methods

Each Go SDK method used in the Atlas CLI should be marked as stable.
Stable methods are listed in the SDK GO [operations.stable.json](https://github.com/mongodb/atlas-sdk-go/blob/main/tools/transformer/src/operations.stable.json) file.

We have developed automation that lists stable methods.
Generate list from Atlas CLI run:

```
go run ./tools/sdk-usage/main.go ./internal/store ./operations.stable.json
```

After file is create please create PR directly in the GO SDK containing updated file.

in order to update `operations.stable.json` file in the Go SDK.
