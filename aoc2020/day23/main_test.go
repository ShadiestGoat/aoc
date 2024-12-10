package day23_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day23"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

var INPUT = `389125467`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day23.Solve1, `67384529`)
}
