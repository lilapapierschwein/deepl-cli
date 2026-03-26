# deepl-cli

a command line utility to interact with deepl (web)

## installation

```sh
# clone the repo
git clone https://github.com/lilapapierschwein/deepl-cli.git $INSTALL_DIRECTORY

# cd in to directory and run installer
cd $INSTALL_DIRECTORY && ./install.sh
```

## usage

- after installation run `deepl --help`
- shell completion can be generated with `deepl completion $YOUR_SHELL` for supported shells:
  - bash
  - fish
  - powershell
  - zsh

- use

## motivation

`deepl-cli` is a fast and lightweight command line interface that utilizes [deepl.com](https://deepl.com)'s web app for translations.

while eliminating the need for an api-key, this comes with a set of restrictions on usage:

- no account: 1.5k characters per translation
- free account: 5k characters per translation
- pro account: 20k characters per translation

despite those restrictions, the goal is to provide an easy-to-use, data-efficient translation tool for short- to mid-sized text.
