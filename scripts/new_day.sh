#!/bin/sh

DAY="$(/bin/ls -d ./aoc2024/day* | wc -w)"
DAY="$(($DAY + 1))"
YEAR="$(cat .conf/.year)"
YEAR_DIR="$(dirname $0)/../aoc$YEAR"
DAY_DIR="$YEAR_DIR/day$DAY"

if [[ -d $DAY_DIR ]]; then
    echo "[ERR] Grr... Dir exits"
    exit 1
fi

mkdir "$DAY_DIR"

cat > "$DAY_DIR/main.go" <<EOF
package day$DAY

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

cat > "$DAY_DIR/main_test.go" <<EOF
package day${DAY}_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc$YEAR/day$DAY"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = \`

\`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day$DAY.Solve1, VALUE)
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day$DAY.Solve2, VALUE)
//}
EOF
