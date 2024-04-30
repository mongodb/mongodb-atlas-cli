# MongoDB CLI

The MongoDB CLI is a modern command line interface that enables you to manage your MongoDB services from the terminal.

![mongocli-atlas-quickstart](https://user-images.githubusercontent.com/461027/126986233-0dd5c82a-2c75-4887-ab66-eb018c59e093.gif)

Use simple, one-line commands to interact with MongoDB Atlas, Cloud Manager, or Ops Manager, and automate management tasks for your deployments.

## Documentation

See the [official docs](https://docs.mongodb.com/mongocli/stable/) for instructions on how to
install, set up, and reference available commands.

## Installing

### Homebrew on macOS

```bash
brew install mongocli
```

### Pre-built Binaries

Download the appropriate version for your platform from [mongocli releases](https://github.com/mongodb/mongodb-atlas-cli/releases).
After you download the library, you can run it from anywhere and don't need to install it into a global location.
This works well for shared hosts and other systems where you don't have a privileged account.

You can place this binary somewhere in your `PATH` for ease of use.
`/usr/local/bin` is the most probable location.

### Build From Source

#### Fetch Source

```bash
git clone https://github.com/mongodb/mongodb-atlas-cli.git
cd mongodb-atlas-cli
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

To get a list of available commands, run `mongocli help`
or check our [documentation](https://docs.mongodb.com/mongocli/master/) for more details.

### Configuring MongoCLI with Atlas
To use `mongocli` with Atlas, open your terminal, run `mongocli auth login`, and follow the prompted steps.

### Configuring MongoCLI with Ops Manager and Cloud Manager

#### Getting API Keys (Ops manager / Cloud Manager)
To use `mongocli`, create API keys. To learn more, see the documentation for the service you're using:
- [Atlas](https://docs.atlas.mongodb.com/configure-api-access/),
- [Ops Manager](https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/),
- [Cloud Manager](https://docs.cloudmanager.mongodb.com/tutorial/manage-programmatic-api-keys/)

#### Set up your credentials
To set up your credentials, run `mongocli config`, or use [env variables](https://docs.mongodb.com/mongocli/stable/configure/environment-variables/) instead.

If you're working with Atlas Gov, Ops Manager or Cloud Manager you need to define the service using `--service`

- For Atlas Gov, `mongocli config --service cloudgov`
- For Ops Manager, `mongocli config --service ops-manager`
- For Cloud Manager, `mongocli config --service cloud-manager`

### Shell Completions

If you install via [homebrew](#hombrew-on-macos) no additional actions are needed.

To get specific instructions for your preferred shell, run:

```bash
mongocli completion <bash|zsh|fish|powershell> --help
```