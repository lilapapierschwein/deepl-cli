update_deepl_cli() {
    if [[ $EUID -ne 0 ]]; then
        echo "script must be run as root!"
        return 1
    fi

    local GO_BIN=/usr/local/go/bin/go
    if [[ ! -x $GO_BIN ]]; then
        echo "go executable not found at ${GO_BIN%%/bin/go} (${GO_BIN%%/go}). cancelled."
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

    local PROJECT_BASE_NAME="deepl-cli"
    local MAIN_FILE_NAME="deepl.go"

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

    if [[ ! -x $BIN_PATH ]]; then
        echo "error: binary not found at $BIN_PATH"
        echo "it seems the application has not yet been installed."
        echo ""
        echo "*to build and install, run the script at ./run/install.sh."
        return 1
    fi

    echo "updating deepl-cli..."
    echo ""

    cd $PROJECT_PATH

    echo "rebuilding binary from file $PROJECT_PATH/$MAIN_FILE_NAME..."
    $GO_BIN build $MAIN_FILE_NAME
    echo "binary rebuilt at $PROJECT_PATH/$BIN_NAME"
    echo ""

    echo "removing old binary from $BIN_PATH..."
    rm $BIN_PATH
    echo "old binary removed"
    echo ""

    echo "linking updated binary..."
    ln -s $PROJECT_PATH/$BIN_NAME $BIN_PATH
    echo "updated binary linked to $BIN_PATH"
    echo ""

    echo "the application was updated successfully"
    echo "run \`deepl --help\` for info on usage and shell completions"
}

update_deepl_cli $@
