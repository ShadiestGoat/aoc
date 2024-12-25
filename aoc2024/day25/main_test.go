package day25_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day25"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day25.Solve1, 3)
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day25.Solve2, VALUE)
//}
