package day17_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day17"
)

const INPUT = `
.#.
..#
###
`

func TestSolve1(t *testing.T) {
	ans := day17.Solve1(strings.TrimSpace(INPUT))

	if ans != 112 {
		t.Fatalf(":( expected %v, got %v", 112, ans)
	}
}
