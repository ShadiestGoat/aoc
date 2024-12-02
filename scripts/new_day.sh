#!/bin/sh

if [[ -z $1 ]]; then
    echo "[ERR] Grr... Wheres the day?"
    exit 1
fi

YEAR="$(cat .conf/.year)"
YEAR_DIR="$(dirname $0)/../aoc$YEAR"
DAY_DIR="$YEAR_DIR/day$1"

if [[ -d $DAY_DIR ]]; then
    echo "[ERR] Grr... Dir exits"
    exit 1
fi

mkdir "$DAY_DIR"

cat > "$DAY_DIR/main.go" <<EOF
package day$1

func parseInput(inp string) any {
	return nil
}

func Solve1(inp string) any {
	return nil
}

func Solve2(inp string) any {
	return nil
}
EOF
