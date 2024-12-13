#!/bin/sh

YEAR="$(cat .conf/.year)"

go test ./aoc$YEAR/...