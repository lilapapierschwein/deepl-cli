install_test() {
    local INSTALL_PATH=/usr/local/bin
    local BIN_PATH=$INSTALL_PATH/deepl
    local PROJECT_BASE_NAME="deepl-cli"

    local PROJECT_PATH=$PWD
    while true; do
        if [[ "$(basename $PROJECT_PATH)" != "$PROJECT_BASE_NAME" ]]; then
            PROJECT_PATH="${PROJECT_PATH%/*}"
        else
            PROJECT_PATH=$(realpath $PROJECT_PATH)
            break
        fi
    done

    echo $PROJECT_PATH
}

install_test
