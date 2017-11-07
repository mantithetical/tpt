#!/usr/bin/env bash

# Supports Mac OSs only

if [[ "$OSTYPE" == "darwin"* ]]; then
    pkg="$GOPATH/pkg/darwin_amd64/github.com/mantithetical/tpt"
    echo "Deleting $pkg..."
    rm -fr $pkg;
    bin="$GOPATH/bin/tpt...";
    echo "Deleting $bin";
    rm -f $bin;
else
    echo "$OSTYPE is not supported";
fi