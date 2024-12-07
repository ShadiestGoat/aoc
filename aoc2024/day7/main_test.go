package day7_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day7"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day7.Solve1, 3749)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day7.Solve2, 11387)
}
