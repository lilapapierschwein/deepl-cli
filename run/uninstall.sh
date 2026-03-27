uninstall() {
    if [[ $EUID -ne 0 ]]; then
        echo "script must be run as root!"
        return 1
    fi

    local INSTALL_PATH=/usr/local/bin

    local BIN_NAME="deepl"
    local BIN_PATH=$INSTALL_PATH/$BIN_NAME

    if [[ ! -x $BIN_PATH ]]; then
        echo "error: binary not found at '$BIN_PATH'"
        echo ""
        echo "*to install the application, run the installer at ./run/install.sh"
        return 1
    fi

    rm $BIN_PATH

    echo "application successfully uninstalled (binary removed) from $BIN_PATH"
}

uninstall
