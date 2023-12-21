#!/bin/sh

# fatal: detected dubious ownership in repository at '...'
git config --global --add safe.directory "$(pwd)"
CGO_ENABLED=0 go build -tags timetzdata --ldflags "-X 'main.Version=$(git describe --tags)' -extldflags \"-static\" -s -w" -o sak .
