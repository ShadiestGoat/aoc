package day25_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day25"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
5764801
17807724
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day25.Solve1, 14897079)
}
