#!/usr/bin/env bash

install_deepl_cli() {
    if [[ $EUID -ne 0 ]]; then
        echo "script must be run as root!"
        return 1
    fi

    local GO_BIN=/usr/local/go/bin/go
    if [[ ! -x $GO_BIN ]]; then
        echo "go executable not found at ${GO_BIN%%/bin/go} (${GO_BIN%%/go}). cancelled."
        return 1
    fi

    local INSTALL_PATH=/usr/local/bin

    local BIN_NAME="deepl"
    local BIN_PATH=$INSTALL_PATH/$BIN_NAME

    local PROJECT_BASE_NAME="deepl-cli"
    local MAIN_FILE_NAME=deepl.go

    local PROJECT_PATH=$PWD
    while true; do
        if [[ "$(basename $PROJECT_PATH)" != "$PROJECT_BASE_NAME" ]]; then
            PROJECT_PATH="${PROJECT_PATH%/*}"
        elif [[ "$(basename $PROJECT_PATH)" = "home" ]]; then
            echo "unable to find project's base directory ($PROJECT_BASE_NAME)."
            echo "operation cancelled."
            return 1
        else
            PROJECT_PATH=$(realpath $PROJECT_PATH)
            break
        fi
    done

    if [[ -x "$BIN_PATH" ]]; then
        echo "error: the application is already installed at $BIN_PATH"
        echo ""
        echo "*to update, run the updater at ./run/update.sh"
        echo "*to uninstall, run the uninstall script at ./run/uninstall.sh"
        return 1
    fi

    echo "installing deepl-cli..."
    echo ""

    cd $PROJECT_PATH

    echo "building binary from file $PROJECT_PATH/$MAIN_FILE_NAME..."
    $GO_BIN build $MAIN_FILE_NAME
    echo "binary built at $PROJECT_PATH/$BIN_NAME"
    echo ""

    echo "linking binary..."
    ln -s $PROJECT_PATH/$BIN_NAME $BIN_PATH
    echo "binary linked to $BIN_PATH"
    echo ""

    echo "application successfully installed at $BIN_PATH."
    echo "make sure $INSTALL_PATH is in \$PATH. to add run:"
    echo "  \`export PATH=\$PATH:$INSTALL_PATH\`"
    echo ""
    echo "run \`deepl --help\` for info on usage and shell completions"
}

install_deepl_cli
