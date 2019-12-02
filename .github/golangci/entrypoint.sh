#!/bin/sh

set -e

APP_DIR="$GOPATH/src/github.com/${GITHUB_REPOSITORY}"

mkdir -p ${APP_DIR}
cp -r ./ ${APP_DIR} && cd ${APP_DIR}

echo "APP_DIR" $APP_DIR

echo "Run Golangci-lint"
golangci-lint run --skip-dirs=vendor

# echo "Go Test"
# go test -v -cover -race ./...