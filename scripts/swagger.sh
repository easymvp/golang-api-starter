#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE[0]}")/../" || exit
swag init -g cmd/main.go -o ./docs
