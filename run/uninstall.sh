#!/bin/bash

# check for root/sudo
if [[ $EUID -ne 0 ]]; then
    echo "script must be run as root!"
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

if [[ ! -x $bin_path ]]; then
    echo "error: binary not found at '$bin_path'"
    echo ""
    echo "*to install the application, run the installer at ./run/install.sh"
    exit 1
fi

rm $bin_path

echo "application successfully uninstalled (binary removed) from $bin_path"

exit 0
