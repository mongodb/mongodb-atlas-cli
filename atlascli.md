## AtlasCLI (Pre-Release)
![GO tests](https://github.com/mongodb/mongocli/workflows/GO%20tests/badge.svg)
![golangci-lint](https://github.com/mongodb/mongocli/workflows/golangci-lint/badge.svg)

The MongoDB AtlasCLI is a modern command line interface that enables you to manage MongoDB Atlas from the terminal.

![atlascli-atlas-quickstart](https://user-images.githubusercontent.com/5663078/156184669-57c8ddce-6f0a-4e84-9311-2d996cb27942.gif)

## Installing

AtlasCLI is currently in the pre-release phase, do not use it in the production environment.

### Pre-built Binaries

Download the appropriate version for your platform from [AtlasCLI releases](https://github.com/mongodb/mongocli/releases).
After you download the library, you can run it from anywhere and don't need to install it into a global location.
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should place this binary somewhere in your `PATH` for ease of use.
`/usr/local/bin` is the most probable location.

### Build From Source

#### Fetch Source

```bash
git clone https://github.com/mongodb/mongocli.git
cd mongocli
```

#### Build

To build `atlascli`, run:

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

To get a list of available commands, run `atlas help`
or check our online documentation for more details.

### Configuring AtlasCLI
To use `atlascli` with Atlas, open your terminal, run `atlas auth login`, and follow the prompted steps.