#!/bin/sh

YEAR="$(cat .conf/.year)"

go generate ./...
go run -tags "$YEAR" -ldflags="-X 'github.com/shadiestgoat/aoc/solvers.YEAR=$YEAR'" . ${@:1}
