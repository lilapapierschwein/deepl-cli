install_deepl_cli () {
    echo "running installer for deepl-cli." 
    echo "this will downlad the application's binary onto your system."
    echo ""

    read -p "Proceed? [Y/n]: " USER_OKAY
    echo ""

    if [[ $USER_OKAY = "n" || $USER_OKAY = "no" || $USER_OKAY = "N" ]]; then
        echo "installation cancelled."
        return 0
    fi
    
    curl -O https://github.com/lilapapierschwein/deepl-cli/raw/refs/heads/installer/bin/deepl
    mv ~/Downloads/deepl /usr/local/bin

    echo "application successfully installed at $install_path/deepl"
    echo "run \`deepl --help\` for info on usage and shell completions"
    
    return 0
}

install_deepl_cli
