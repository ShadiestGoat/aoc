package day3_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day3"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day3.Solve1, 159)
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day3.Solve2, VALUE)
//}
