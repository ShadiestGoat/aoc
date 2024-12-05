package day4_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day4"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day4.Solve1, 18)
}
