install () {
    local install_path=/usr/local/bin/
    
    go build main.go
    mv main deepl
    ln -s $PWD/deepl $install_path
}

install
