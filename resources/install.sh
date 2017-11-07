#!/usr/bin/env bash

brew install go;

if [ -z "$GOPATH" ]; then
    mkdir $HOME/go;
    export GOPATH=$HOME/go;
fi

go get github.com/mantithetical/tpt;
cd $GOPATH/src/github.com/mantithetical/tpt;
export PATH=$PATH:$GOPATH/bin

go get ./...;
go install;

cp resources/.tpt.yaml $HOME/.;

brew install bash-completion;

cp resources/tpt_bash_completion.sh $(brew --prefix)/etc/bash_completion.d/tpt_bash_completion.sh;

if [ -f $(brew --prefix)/etc/bash_completion ]; then
. $(brew --prefix)/etc/bash_completion
fi
