#!/bin/bash

# check for root/sudo
if [[ $EUID -ne 0 ]]; then
    echo "script must be run as root!"
    exit 1
fi

# verify go installation & get binary
go_bin=/usr//go/bin/go
if [[ ! -x $go_bin ]]; then
    echo "go executable not found at ${go_bin%%/bin/go} (${go_bin%%/go}). cancelled."
    exit 1
fi

default_install_path=/usr//bin
install_path="$1"

if [[ -z $install_path ]]; then
    install_path=$(realpath $default_install_path)
else
    install_path=$(realpath $install_path)
fi

bin_name="deepl"
bin_path=$install_path/$bin_name

project_base_name="deepl-cli"
main_file_name=deepl.go

project_path=$PWD
while true; do
    if [[ "$(basename $project_path)" != "$project_base_name" ]]; then
        project_path="${project_path%/*}"
    elif [[ "$(basename $project_path)" = "home" ]]; then
        echo "unable to find project's base directory ($project_base_name)."
        echo "operation canceled."
        exit 1
    else
        project_path=$(realpath $project_path)
        break
    fi
done

if [[ -x "$bin_path" ]]; then
    echo "error: the application is already installed at $bin_path"
    echo ""
    echo "*to update, run the updater at ./run/update.sh"
    echo "*to uninstall, run the uninstall script at ./run/uninstall.sh"
    exit 1
fi

echo "installing deepl-cli..."
echo ""

cd $project_path

echo "building binary from file $project_path/$main_file_name..."
$go_bin build $main_file_name
echo "binary built at $project_path/$bin_name"
echo ""

echo "linking binary..."
ln -s $project_path/$bin_name $bin_path
echo "binary linked to $bin_path"
echo ""

echo "application successfully installed at $bin_path."
echo "make sure $install_path is in \$PATH. to add run:"
echo "  \`export PATH=\$PATH:$install_path\`"
echo ""
echo "run \`deepl --help\` for info on usage and shell completions"
