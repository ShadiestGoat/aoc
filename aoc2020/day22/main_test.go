package day22_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day22"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day22.Solve1, 306)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day22.Solve2, 291)
}