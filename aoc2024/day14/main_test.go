package day14_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day14"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`

func TestGenericSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, func(s string) any {
		return day14.GenericSolve1(s, 100, 11, 7)
	}, 12)
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day14.Solve2, VALUE)
//}
