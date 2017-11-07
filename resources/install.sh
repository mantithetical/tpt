#!/usr/bin/env bash

brew install go;

go get github.com/mantithetical/tpt;
cd $GOPATH/src/github.com/mantithetical/tpt;
got get ./...;
go install;

cp resources/.tpt.yaml $HOME/.;

if brew ls --versions bash-completion > /dev/null; then
    brew install bash-completion;
fi

cp resources/tpt_bash_completion.sh $(brew --prefix)/etc/bash_completion.d/tpt_bash_completion.sh;

touch $HOME/.bashrc;
cat <<EOT >> $HOME/.bashrc
if [ -f $(brew --prefix)/etc/bash_completion ]; then
. $(brew --prefix)/etc/bash_completion
fi
EOT
