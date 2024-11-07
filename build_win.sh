#!/bin/bash
set -e
export https_proxy=http://127.0.0.1:1087 http_proxy=http://127.0.0.1:1087 all_proxy=socks5://127.0.0.1:1087

export CGO_ENABLED=0
export GO111MODULE=on
export GOOS=windows
export GOARCH=amd64

export target=build/wrest.exe

echo "building for ${GOOS}/${GOARCH}"

go build -ldflags="-s -w" -o ${target} main.go