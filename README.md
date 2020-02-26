# MongoDB CLI
[![Build Status](https://cloud.drone.io/api/badges/mongodb/mongocli/status.svg)](https://cloud.drone.io/mongodb/mongocli)

`mongocli` is a tool for managing your MongoDB cloud services

![Screenshot 2020-01-03 at 10 49 27](https://user-images.githubusercontent.com/461027/73472371-d6656480-4382-11ea-8bfb-535d91cf1fb4.png)

# Installing

## Hombrew on macOS

```bash
brew install mongodb-labs/mongocli/mongocli
```

## Pre-built Binaries
Download the appropriate version for your platform from [mongocli releases](https://github.com/mongodb/mongocli/releases). 
Once downloaded, the binary can be run from anywhere.
You don't need to install it into a global location. 
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should install it somewhere in your PATH for easy use. `/usr/local/bin` is the most probable location.

## Build from Source 

### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.13)](https://golang.org/dl/)

### Fetch and Install

```bash
git clone https://github.com/mongodb/mongocli.git
cd mongocli
make install
```

# Usage

Run `mongocli help` for a list of available commands

## Authentication
Run `mongocli config` to set up a profile.

You can also use `MCLI_OPS_MANAGER_URL`, `MCLI_PUBLIC_API_KEY`, and `MCLI_PRIVATE_API_KEY` to define some of the authentication variables

## ZSH Completion (experimental)
Add the following to your `.zshrc` file

```bash
source <(mongocli completion zsh)
compdef _mongocli mongocli
```

# Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

# License

MongoDB CLI is released under the Apache 2.0 license. See [LICENSE](LICENSE)
