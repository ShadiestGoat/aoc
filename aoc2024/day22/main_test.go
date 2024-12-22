package day22_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day22"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
1
10
100
2024
`

const INPUT_2 = `
1
2
3
2024
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT_1, day22.Solve1, 37327623)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT_2, day22.Solve2, 23)
}
