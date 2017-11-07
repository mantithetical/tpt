#!/usr/bin/env bash

# Supports Mac OSs only

declare -a deps=("c9s" "fsnotify" "jdkato" "magiconair" "mitchellh" "pelletier"
"spf13")

if [ -z "$GOPATH" ]; then
    export GOPATH=$HOME/go;
fi
export PATH=$PATH:$GOPATH/bin

if [[ "$OSTYPE" == "darwin"* ]]; then
    for i in "${deps[@]}"
    do
        pkg="$GOPATH/pkg/darwin_amd64/github.com/$i"
        src="$GOPATH/src/github.com/$i"
        echo "Deleting $pkg..."
        rm -fr $pkg;
        echo "Deleting $src..."
        rm -fr $src;
    done

    bin="$GOPATH/bin/tpt";
    rm -f $bin;
    rm -f $HOME/.tpt.yaml
    rm -f $(brew --prefix)/etc/bash_completion.d/tpt_bash_completion.sh
else
    echo "$OSTYPE is not supported";
fi