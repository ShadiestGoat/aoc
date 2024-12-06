package day6_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day6"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day6.Solve1, 41)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day6.Solve2, 6)
}
