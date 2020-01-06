# MCLI

`mcli` is a tool for managing your MongoDB cloud services

![Screenshot 2020-01-03 at 10 49 27](https://user-images.githubusercontent.com/461027/71719742-2e0dc000-2e17-11ea-885c-385a80aea95a.png)

# Installing
## Binary
Download the appropriate version for your platform from [mcli releases](https://github.com/10gen/mcli/releases). 
Once downloaded, the binary can be run from anywhere.
You don't need to install it into a global location. 
This works well for shared hosts and other systems where you don't have a privileged account.

Ideally, you should install it somewhere in your PATH for easy use. `/usr/local/bin` is the most probable location.

## Source 

### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.13)](https://golang.org/dl/)

## Fetch from GitHub 
The easiest way to get started is to clone `mcli` and install with go:

```bash
git clone https://github.com/10gen/mcli.git
cd mcli
make install
```

# Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

# License

[SSPL](LICENSE)
