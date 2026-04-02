#!/bin/bash

# check for root/sudo
if [[ $EUID -ne 0 ]]; then
    echo "script must be run as root!"
    exit 1
fi

# verify go installation & get binary
go_bin=/usr/local/go/bin/go
if [[ ! -x $go_bin ]]; then
    echo "go executable not found at ${go_bin%%/bin/go} (${go_bin%%/go}). cancelled."
    exit 1
fi

default_install_path=/usr/local/bin
install_path="$1"

if [[ -z $install_path ]]; then
    install_path=$(realpath $default_install_path)
else
    install_path=$(realpath $install_path)
fi

bin_name="deepl"
bin_path=$install_path/$bin_name

project_base_name="deepl-cli"
main_file_name="deepl.go"

project_path=$PWD
while true; do
    if [[ "$(basename $project_path)" != "$project_base_name" ]]; then
        project_path="${project_path%/*}"
    elif [[ "$(basename $project_path)" = "home" ]]; then
        echo "unable to find project's base directory ($project_base_name)."
        echo "operation cancelled."
        exit 1
    else
        project_path=$(realpath $project_path)
        break
    fi
done

if [[ ! -x $bin_path ]]; then
    echo "error: binary not found at $bin_path"
    echo "it seems the application has not yet been installed."
    echo ""
    echo "*to build and install, run the script at ./run/install.sh."
    exit 1
fi

echo "starting deepl-cli update..."
echo ""

cd $project_path

# check remote origin for updates
branch="main"
echo "checking for updates on '$branch' branch..."

git fetch origin $branch >/dev/null 2>&1
git_status="$(git status)"
git_status_summary="$(echo "$git_status" | sed -n '2 p')"

if [[ "$git_status_summary" = "Your branch is up to date with 'origin/$branch'." ]]; then
    echo "no updates found. '$branch' is up to date."
    exit 0
elif [[ "$git_status_summary" =~ ^((Your branch is behind.+/)([A-Za-z]+)\' by ([0-9]+).+(can be fast-forwarded){1}.)$ ]]; then
    updates_count="${BASH_REMATCH[4]}"
    branch_verified="${BASH_REMATCH[3]}"
    fast_forward="${BASH_REMATCH[5]}"

    echo "$updates_count updates found on '$branch_verified' branch. starting update..."
else
    echo "WTF?! THIS IS TEST!"
    exit 1
fi

echo "pulling changes from origin/$branch_verified..."
git pull origin $branch >/dev/null 2>&1
echo "download complete"

echo "rebuilding binary from file $project_path/$main_file_name..."
$go_bin build $main_file_name
echo "binary rebuilt at $project_path/$bin_name"
echo ""

echo "removing old binary from $bin_path..."
rm $bin_path
echo "old binary removed"
echo ""

echo "linking updated binary..."
ln -s $project_path/$bin_name $bin_path
echo "updated binary linked to $bin_path"
echo ""

echo "the application was updated successfully"
echo "run \`deepl --help\` for info on usage and shell completions"

exit 0
