## MongoDB CLI

![GO tests](https://github.com/mongodb/mongocli/workflows/GO%20tests/badge.svg)
![golangci-lint](https://github.com/mongodb/mongocli/workflows/golangci-lint/badge.svg)

The MongoDB CLI is a modern command line interface that enables you to manage your MongoDB services from the terminal.

![mongocli-atlas-quickstart](https://user-images.githubusercontent.com/461027/126986233-0dd5c82a-2c75-4887-ab66-eb018c59e093.gif)

Use simple, one-line commands to interact with MongoDB Atlas, Cloud Manager, or Ops Manager, and to automate management tasks for your deployments.

## Documentation

See the [official docs](https://docs.mongodb.com/mongocli/stable/) for instructions on how to
install, set up, and reference available commands.

## Installing

### Homebrew on macOS

```bash
brew install mongocli
```

### Pre-built Binaries

Download the appropriate version for your platform from [mongocli releases](https://github.com/mongodb/mongocli/releases). 
Once downloaded, the binary can be run from anywhere.
You don't need to install it into a global location. 
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should place this binary somewhere in your `PATH` for easy use. 
`/usr/local/bin` is the most probable location.

### Build From Source 

#### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.17)](https://golang.org/dl/)

#### Fetch Source

```bash
git clone https://github.com/mongodb/mongocli.git
cd mongocli
```

#### Build

To build `mongocli`, run:

```bash
make build
```

The resulting `mongocli` binary is placed in `./bin`.

#### Install

To install the `mongocli` binary in `$GOPATH/bin`, run:

```bash
make install
```

**Note:** running `make build` is not needed when running `make install`.

## Usage

Run `mongocli help` for a list of available commands
or check our [online documentation](https://docs.mongodb.com/mongocli/master/) for more details.

### Configuring `mongocli` with Atlas
To use `mongocli` with Atlas you'll just need to run `mongocli auth login` and follow the prompted steps on your terminal.

### Configuring `mongocli` with Ops Manager and Cloud Manager

#### Getting API Keys (Ops manager / Cloud Manager)
To use `mongocli` you'll need to get API keys, to get them follow the documentation
appropriate for the service you're using,
[Atlas](https://docs.atlas.mongodb.com/configure-api-access/),
[Ops Manager](https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/),
or [Cloud Manager](https://docs.cloudmanager.mongodb.com/tutorial/manage-programmatic-api-keys/)

#### Set up your credentials
Run `mongocli config` to set up your credentials, 
this is optional, you can use [env variables](https://docs.mongodb.com/mongocli/stable/configure/environment-variables/) instead.

If you're working with Atlas Gov, Ops Manager or Cloud Manager you need to define the service using `--service`

For Atlas Gov, `mongocli config --service cloudgov`.

For Ops Manager, `mongocli config --service ops-manager`.

For Cloud Manager, `mongocli config --service cloud-manager`.

### Shell Completions

If you install via [homebrew](#hombrew-on-macos) no additional actions are needed.

To get specific instructions for your preferred shell run:

```bash
mongocli completion <bash|zsh|fish|powershell> --help
```

## Atlas CLI (Pre-Release)
![GO tests](https://github.com/mongodb/mongocli/workflows/GO%20tests/badge.svg)
![golangci-lint](https://github.com/mongodb/mongocli/workflows/golangci-lint/badge.svg)

The MongoDB Atlas CLI is a modern command line interface that enables you to manage MongoDB Atlas from the terminal.

## Installing

Atlas CLI is currently in the pre-release phase, so it should not be used in production environment. 

### Pre-built Binaries

Download the appropriate version for your platform from [Atlas CLI releases](https://github.com/mongodb/mongocli/releases).
Once downloaded, the binary can be run from anywhere.
You don't need to install it into a global location.
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should place this binary somewhere in your `PATH` for easy use.
`/usr/local/bin` is the most probable location.

### Build From Source

#### Prerequisite Tools
- [Git](https://git-scm.com/)
- [Go (at least Go 1.17)](https://golang.org/dl/)

#### Fetch Source

```bash
git clone https://github.com/mongodb/mongocli.git
cd mongocli
```

#### Build

To build `Atlas CLI`, run:

```bash
make build-atlascli
```

The resulting `atlas` binary is placed in `./bin`.

#### Install

To install the `atlas` binary in `$GOPATH/bin`, run:

```bash
make install-atlascli
```

**Note:** running `make build-atlascli` is not needed when running `make install-atlascli`.


## Usage

Run `atlas help` for a list of available commands
or check our online documentation for more details.

### Configuring `Atlas CLI`
To use `Atlas CLI` you'll just need to run `atlas auth login` and follow the prompted steps on your terminal.


## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) guide.

## License

MongoDB CLI and Atlas CLI are released under the Apache 2.0 license. See [LICENSE](LICENSE)
