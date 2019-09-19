semver-git-tag
===

Simple helper to release your git tag according to the semver spec

## Installation
### using go ...
``go get -u github.com/timo-reymann/git-semver-tag``

### using Linux ...
```bash
curl -LO https://github.com/timo-reymann/git-semver-tag/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/git-semver-tag/releases/latest | grep -o '[^/]*$')/git-semver-tag_linux_amd64
chmod +x git-semver-tag_linux_amd64
sudo mv git-semver-tag_linux_amd64 /usr/local/bin/git-semver-tag
```

### using Mac ...
```bash
curl -LO https://github.com/timo-reymann/git-semver-tag/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/git-semver-tag/releases/latest | grep -o '[^/]*$')/git-semver-tag_linux_amd64
chmod +x git-semver-tag_darwin_amd64
sudo mv git-semver-tag_darwin_amd64 /usr/local/bin/git-semver-tag
```

### using Windows ...
1. Download the latest binary from the [release page](https://github.com/timo-reymann/git-semver-tag/releases)
2. Add the location of the binary to your PATH
3. You are ready to go!

**ATTENTION: Windows support is available, but no support for windows or any guarantee is given!**

## Usage
To get usage info enter:

`git semver-tag --help`

As you may have recognized now semver-tag is registered as subdomain.