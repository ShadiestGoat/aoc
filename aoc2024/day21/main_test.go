package day21_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day21"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
029A
980A
179A
456A
379A
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day21.Solve1, 126384)
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day21.Solve2, VALUE)
//}
