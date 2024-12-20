package day20_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day20"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`

func TestGenericSolve(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		tutils.AssertFunc(t, INPUT, func(s string) any {
			return day20.GenericSolve(s, 2, 1)
		}, 44)
	})

	t.Run("2", func(t *testing.T) {
		tutils.AssertFunc(t, INPUT, func(s string) any {
			return day20.GenericSolve(s, 20, 50)
		}, 285)
	})
}
