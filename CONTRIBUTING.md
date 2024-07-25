# Contributing to MongoDB Atlas CLI

Thanks for your interest in contributing to MongoDB Atlas CLI,
this document describes some guidelines necessary to participate in the community.

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
- [Go (at least Go 1.22)](https://golang.org/dl/)

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
- Run `E2E_TAGS=e2e,atlas make e2e-test` will run end-to-end tests against an Atlas instance,
  please make sure to have set `MCLI_*` variables pointing to that instance.
- Run `E2E_TAGS=cloudmanager,remote,generic make e2e-test` will run end-to-end tests against a Cloud Manager instance.<br />
  Please remember to: (a) have a running automation agent, and (b) set MCLI\_\* variables to point to your Cloud Manager instance.
- Run `make build` to generate a local binary in the `./bin` folder.

We provide a git pre-commit hook to format and check the code, to install it run `make link-git-hooks`.

#### Generating Mocks

We use [mockgen](https://github.com/golang/mock) to handle mocking in our unit tests.
If you need a new mock please update or add the `//go:generate` instruction to the appropriate file.

#### Compilation in VSCode

Please add the following line to your settings.json file :
```
    "go.buildTags": "unit,e2e",
    "go.testTags": "unit,e2e"
```

This will enable compilation for unit test and end-to-end tests.

#### Debugging in VSCode

To debut in VSCode you need to create a debug configuration for the command with required arguments.
Run the following commands to 

```
touch .vscode/launch.json
```
Then put the following configuration into the file.
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
            "args": [
              "login"
            ]
      }
    ]
} 

```

### Contributing New Command Group

`Atlas CLI` uses the [Cobra Framework](https://umarcor.github.io/cobra/).

Depending on the feature you are building you might choose to:

- Add individual commands to existing groups of commands
- Add a new command group that provides ability to run nested commands under certain prefix

For a command group, we need to create new cobra root command. 
This command aggregates a number of subcommands that can perform network requests and return results

For example [teams](https://github.com/mongodb/mongodb-atlas-cli/tree/220c6c73f346f5c711a1c772b17f93a6811efc69/internal/cli/atlas/teams) 
command root provides the main execution point for `atlas teams` with subcommands like `atlas teams list`

Root command links to a number of child commands. Atlas CLI provides a number of patterns for child commands depending on the type of operation performed.
Each new feature might cover typical commands like `list` and `describe` along with dedicated actions.
For example [list command](https://github.com/mongodb/mongodb-atlas-cli/blob/220c6c73f346f5c711a1c772b17f93a6811efc69/internal/cli/atlas/teams/list.go).
It is normal to duplicate existing commands and edit descriptions and methods for your own needs.

Additionally, after adding new command we need to add it to the main CLI root command. 
For example please edit `./root/atlas/builder.go` to add your command builder method for Atlas CLI

### Adding a New Command

`atlas` has defined a basic structure for individual commands that should be followed.
For a `atlas scope newCommand` command, a file `internal/cli/scope/new_command.go` should implement:

- A `ScopeNewCommandOpts` struct which handles the different options for the command.
- At least a `func (opts *ScopeNewCommandOpts) Run() error` function with the main command logic.
- A `func ScopeNewCommandBuilder() *cobra.Command` function to put together the expected cobra definition along with the `ScopeNewCommandOpts` logic.
- A set of documentation fields further described in the section below.

Commands follow a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) approach to match the APIs, whenever possible.
For that reason, command arguments tend to match the path and query params of the APIs,
with the last param being a required argument and the rest handled via flag options.
For commands that create or modify complex data structures, the use of configuration files is preferred over flag options.

> [!TIP]  
> During the development of the commands we recommend setting `Hidden: true` property to make commands invisible to the end users and documentation.

> [!IMPORTANT]  
> Commands are executing network requests by using `./internal/store` interface that wraps [Atlas Go SDK](https://github.com/mongodb/atlas-sdk-go). 
Before adding a command please make sure that your api exists in the GO SDK. 

> [!TIP]  
> Atlas CLI provides an experimental generator, make sure to try it out in [tools/cli-generator](./tools/cli-generator)

### API Interactions

Atlas CLI use [atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) for all backend integration.
This SDK is updated automatically based on Atlas OpenAPI file.

#### How to define flags:

Flags are a way to modify the command, also may be called "options". Flags always have a long version with two dashes (--state) but may also have a shortcut with one dash and one letter (-s).

`atlas` uses the following types of flags:

- `--flagName value`: this type of flag passes the value to the command. Examples: `--projectId 5efda6aea3f2ed2e7dd6ce05`
- `--booleanFlag`: this flag represents a boolean and it sets the related variable to true when the flag is used, false otherwise. Example: `--force`
- `--flagName value1,value2,..,valueN`: you will also find flags that accept a list of values. This type of flag can be very useful to represent data structures as `--role roleName1@db,roleName2@db`, `--privilege action@dbName.collection,action2@dbName.collection`, or `--key field:type`.
  As shown in the examples, the standard format used to represent data structures consists of splitting the first value with the second one by at sign `@` or colon `:`, and the second value with the third one by a full stop `.`.
  We recommend using configuration files for complex data structures that require more than three values. For an example of configuration files, see [atlas cluster create](https://github.com/mongodb/mongodb-atlas-cli/blob/f2e6d661a3eb2cfcf9baab5f9e0b1c0f872b8c14/internal/cli/atlas/clusters/create.go#L235).

When in doubt refer to [Utility Argument Syntax Conventions](https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html)

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

```sh
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

### Update Automation

To update Atlas SDK run:

```bash
make update-atlas-sdk
```

> [!NOTE]  
> Update mechanism is only needed for major releases. Any other releases will be supported by dependabot.

> [!IMPORTANT]  
> Command can make import changes to +500 files. Please make sure that you perform update on main branch without any uncommited changes.
