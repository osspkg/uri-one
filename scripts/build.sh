#!/bin/bash

PRJROOT="$PWD"
GOMAIN="$PWD/cmd/uri-one"

cd $PWD

rm -rf $PRJROOT/build/bin/uri-one_$1

go generate ./...

# sudo apt install gcc-aarch64-linux-gnu

case $1 in
arm64)
    GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc \
        go build -a -o $PRJROOT/build/bin/uri-one_$1 $GOMAIN
  ;;
front)
  front
  ;;
*)
  GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=$1 go build -o $PRJROOT/build/bin/uri-one_$1 $GOMAIN
  ;;
esac
