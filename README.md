# deepl-cli

a command line utility to interact with deepl (web)

## installation

to install you can choose from 2 options:

1. run the installation script
2. compile from source manually

### via installer

the installation script will download the pre-built binary `deepl` and install it into `/usr/local/bin/`.

```bash
# download installation scipt via `curl` or `wget` and run it
sh -c "$(curl -fsSL https://raw.githubusercontent.com/lilapapierschwein/deepl-cli/refs/heads/main/install.sh)"
sh -c "$(wget -O- https://raw.githubusercontent.com/lilapapierschwein/deepl-cli/refs/heads/main/install.sh)"
```

to uninstall, simply remove the binary with `sudo rm /usr/local/bin/deepl`.

### compile from source

```bash
# create a directory at your desired destination as cd into it
mkdir -p /path/to/deepl-cli
cd /path/to/deepl-cli

# clone the repo
git clone https://github.com/lilapapierschwein/deepl-cli.git .

# build
go build deepl.go

# move or symlink* the binary to a director in your path (might need root privileges):
mv /path/to/deepl-cli/deepl /usr/local/bin
# or
sudo ln -s /path/to/deepl-cli/deepl /usr/local/bin

# ------------------------------------------------------------------------------

# *symlinking the binary might be benifitial to be able to easily update the app in the future:

# get latest updates from git repo and rebuild
git pull origin main
go build deepl.go

# use force option (`-f`) to replace the existing binary
sudo ln -sf /path/to/deepl-cli/deepl /usr/local/bin/deepl
```

## usage

after installation, run `deepl --help` to see available options and commands.

### shell completion

shell completion can be generated with `deepl completion <SHELL>` for supported shells:

- bash
- fish
- powershell
- zsh

it is highly recommended to dynamically create completions on terminal startup via `.bashrc`, `.zshrc` or your teminal's equivalent:

```sh
# example for zsh (put in `.zshrc`)
eval "$(deepl completion zsh)"; compdef _deepl deepl
```

however, you can of course create a completion file in your completions directory, if you prefer:

```sh
deepl completion zsh > /path/to/completions/deepl
```
