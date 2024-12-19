package day18_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day18"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`

func testSolve[T any](g func (string, int, int) T) func (string) any {
	return func(s string) any {
		return g(s, 6, 12)
	}
}

func TestGenericSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, testSolve(day18.GenericSolve1), 22)
}

func TestGenericSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, testSolve(day18.GenericSolve2), "6,1")
}
