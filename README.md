git-semver-tag
===
[![GitHub Release](https://img.shields.io/github/v/release/timo-reymann/git-semver-tag.svg?label=version)](https://github.com/timo-reymann/git-semver-tag/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/git-semver-tag)](https://goreportcard.com/report/github.com/timo-reymann/git-semver-tag)
![GitHub Release status](https://github.com/timo-reymann/git-semver-tag/workflows/GitHub%20Release/badge.svg)
![Continuous Build status](https://github.com/timo-reymann/git-semver-tag/workflows/Continous%20Build/badge.svg)

Simple helper to release your git tag according to the semver spec.

## Features

- create git tags easily following semver guidelines
- prefix-aware, will keep your prefix without any further doings
- feels like a direct integration with git
- ability to add custom suffix to your version tag without having to do it all manually
- push the newly created tag to origin

## Installation

### using go ...

``go get -u github.com/timo-reymann/git-semver-tag``

### on Linux ...

```bash
curl -LO https://github.com/timo-reymann/git-semver-tag/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/git-semver-tag/releases/latest | grep -o '[^/]*$')/git-semver-tag_linux_amd64 && chmod +x git-semver-tag_linux_amd64
sudo mv git-semver-tag_linux_amd64 /usr/local/bin/git-semver-tag
```

### on Mac ...

```bash
curl -LO https://github.com/timo-reymann/git-semver-tag/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/git-semver-tag/releases/latest | grep -o '[^/]*$')/git-semver-tag_darwin_amd64 && chmod +x git-semver-tag_darwin_amd64
sudo mv git-semver-tag_darwin_amd64 /usr/local/bin/git-semver-tag
```

## Usage

To get usage info enter:

`git semver-tag -h`

As you may have recognized now semver-tag is registered as git subcommand. 
