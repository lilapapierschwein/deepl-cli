uninstall() {
    if [[ $EUID -ne 0 ]]; then
        echo "script must be run as root!"
        return 1
    fi

    local DEFAULT_INSTALL_PATH=/usr/local/bin
    local INSTALL_PATH="$1"

    if [[ -z $INSTALL_PATH ]]; then
        INSTALL_PATH=$(realpath $DEFAULT_INSTALL_PATH)
    else
        INSTALL_PATH=$(realpath $INSTALL_PATH)
    fi

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

uninstall $@
