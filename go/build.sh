#! /usr/bin/env sh

DIR=day$1

mkdir -p ./bin
go build -o ./bin/${DIR} -ldflags "-w -s" ./cmd/${DIR}
