package day19_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day19"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day19.Solve1, 6)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day19.Solve2, 16)
}
