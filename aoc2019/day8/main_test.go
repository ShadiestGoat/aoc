package day8_test

import (
	"slices"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day8"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func TestParseInput(t *testing.T) {
	tutils.AssertFuncCustomCompare(t, `123456789012`, func(s string) [][][]int {
		b, _ := day8.ParseInput(s, 3, 2)

		return b
	}, [][][]int{
		{{1, 7}, {2, 8}, {3, 9}},
		{{4, 0}, {5, 1}, {6, 2}},
	}, func(a, b [][][]int) bool {
		return slices.EqualFunc(a, b, func(c, d [][]int) bool {
			return slices.EqualFunc(c, d, func(e, f []int) bool {
				return slices.Equal(e, f)
			})
		})
	})
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day8.Solve2, VALUE)
//}
