package day12_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day12"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`

const INPUT_2 = `
<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`

func TestGenericSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT_1, func(s string) any {
		return day12.GenericSolve1(s, 10)
	}, 179)

	tutils.AssertFunc(t, INPUT_2, func(s string) any {
		return day12.GenericSolve1(s, 100)
	}, 1940)
}

func TestSolve2(t *testing.T) {
	testCases := [][2]any{
		{INPUT_1, 2772},
		{INPUT_2, 4686774924},
	}
	tutils.AssertMany(t, testCases, day12.Solve2)
}
