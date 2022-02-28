#!/bin/bash

PRJROOT="$PWD"
GOMAIN="$PWD/cmd/uri-one"

cd $PWD

rm -rf $PRJROOT/build/bin/uri-one_$1

go generate ./...

GO111MODULE=on GOOS=linux GOARCH=$1 go build -o $PRJROOT/build/bin/uri-one_$1 $GOMAIN

