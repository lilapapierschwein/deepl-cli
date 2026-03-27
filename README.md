# deepl-cli

a command line utility to interact with deepl (web)

## installation

```sh
# clone the repo
mkdir deepl-cli && cd deepl-cli
git clone https://github.com/lilapapierschwein/deepl-cli.git .

# run installer
./run/install.sh
```

the binary is installed to `/usr/local/bin` by default. to alter this behavior, provide your desired destination as an argument to the install script.
same goes for all other scripts in `run` (`update.sh` & `uninstall.sh`).

## usage

- after installation run `deepl --help`
- shell completion can be generated with `deepl completion $YOUR_SHELL` for supported shells:
  - bash
  - fish
  - powershell
  - zsh

it is recommended to source completions in your shell init file (e.g. `.bashrc`, `.zshrc`, ...).
to do this, add the following line to your shell init file:

```sh
# example: zsh
eval "$(deepl completion zsh)"; compdef _deepl deepl
```

## motivation

`deepl-cli` is a fast and lightweight command line interface that utilizes [deepl.com](https://deepl.com)'s web app for translations.

while eliminating the need for an api-key, this comes with a set of restrictions on usage:

- no account: 1.5k characters per translation
- free account: 5k characters per translation
- pro account: 20k characters per translation

despite those restrictions, the goal is to provide an easy-to-use, data-efficient translation tool for short- to mid-sized text.
