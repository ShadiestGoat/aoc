#!/bin/sh

YEAR="$(cat .conf/.year)"

go generate ./...
go run -tags "$YEAR" -ldflags="-X 'shadygoat.eu/aoc/solvers.YEAR=$YEAR'" . ${@:1}
