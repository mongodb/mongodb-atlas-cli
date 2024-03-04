## MongoDB Atlas CLI

The MongoDB Atlas CLI is a modern command line interface that enables you to manage MongoDB Atlas from the terminal.

![atlascli-atlas-quickstart](https://user-images.githubusercontent.com/5663078/156184669-57c8ddce-6f0a-4e84-9311-2d996cb27942.gif)

## Installing

### Pre-built Binaries

Download the appropriate version for your platform from [Atlas CLI releases](https://github.com/mongodb/mongodb-atlas-cli/releases).
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

To build `atlascli`, run:

```bash
make build
```

The resulting `atlas` binary is placed in `./bin`.

#### Install

To install the `atlas` binary in `$GOPATH/bin`, run:

```bash
make install
```

**Note:** running `make build` is not needed when running `make install`.


## Usage

To get a list of available commands, run `atlas help`
or check our documentation for more details.

### Configuring Atlas CLI
To use `atlascli`, open your terminal, run `atlas auth login`, and follow the prompted steps.

### Shell Completions

If you install via [homebrew](#hombrew-on-macos) no additional actions are needed.

To get specific instructions for your preferred shell, run:

```bash
atlas completion <bash|zsh|fish|powershell> --help
```

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) guide.

## License

MongoDB Atlas CLI is released under the Apache 2.0 license. See [LICENSE](LICENSE)
