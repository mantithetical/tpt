#!/usr/bin/env bash

cd $GOPATH/src/github.com/mantithetical/tpt;
got get ./...;
go install;

if brew ls --versions myformula > /dev/null; then
  # The package is installed
fi