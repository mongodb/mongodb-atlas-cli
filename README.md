# MongoDB CLI
[![Build Status](https://cloud.drone.io/api/badges/mongodb/mongocli/status.svg)](https://cloud.drone.io/mongodb/mongocli)

`mongocli` is a tool for managing your MongoDB cloud services

![Screenshot 2020-04-07 at 16 27 14](https://user-images.githubusercontent.com/461027/78688122-d3d83b80-78ec-11ea-84f9-06a24ed7f75a.png)

## Installing

### Hombrew on macOS

```bash
brew tap mongodb/brew
brew install mongocli
```

### Pre-built Binaries
Download the appropriate version for your platform from [mongocli releases](https://github.com/mongodb/mongocli/releases). 
Once downloaded, the binary can be run from anywhere.
You don't need to install it into a global location. 
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should install it somewhere in your PATH for easy use. `/usr/local/bin` is the most probable location.

### Build From Source 

#### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.14)](https://golang.org/dl/)

#### Fetch and Install

```bash
git clone https://github.com/mongodb/mongocli.git
cd mongocli
make install
```

## Usage

Run `mongocli help` for a list of available commands
or check our [online documentation](https://docs.mongodb.com/mongocli/master/) for more details

### Getting Started

Run `mongocli config` to set up your credentials, 
this is optional and you can use [env variables](#environment-variables) instead.

If you're working with Ops Manager or CLoud Manager you need to define the service using `--service`

For Ops Manager, `mongocli config --service ops-manager`.
For Cloud Manager, `mongocli config --service cloud-manager`.  

### Environment Variables

You can use a combination of the next env variables to override your profile settings

- `MCLI_PUBLIC_API_KEY`
- `MCLI_PRIVATE_API_KEY`
- `MCLI_PROJECT_ID`
- `MCLI_ORG_ID`
- `MCLI_OPS_MANAGER_URL`

### ZSH Completion (experimental)
Add the following to your `.zshrc` file

```bash
source <(mongocli completion zsh)
compdef _mongocli mongocli
```

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

MongoDB CLI is released under the Apache 2.0 license. See [LICENSE](LICENSE)
