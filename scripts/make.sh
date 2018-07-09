#!/bin/bash
# Copyright (c) 2018 pirakansa

CMD_DIR=$(cd $(dirname $0);cd ../cmd/file-encrypt/;pwd)
GOBUILD="go build"
GOCLEAN="go clean"
GOTEST="go test"
GOGET="dep ensure"

function __Build() {
    cd ${CMD_DIR}
    ${GOBUILD}
}

function __Clean() {
    cd ${CMD_DIR}
    ${GOCLEAN}
}

function __Install() {
    cd ${CMD_DIR}
    ${GOGET}
}

case "$1" in
    "clean" ) 
        __Clean ;;
    "install" ) 
        __Install ;;
    * ) 
        __Build ;;
esac

