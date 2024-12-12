package day12_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day12"
	"github.com/shadiestgoat/aoc/utils/tutils"
)


const INPUT_1 = `
AAAA
BBCD
BBCC
EEEC
`

const INPUT_2 = `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`

const INPUT_3 = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`

const INPUT_4 = `
00000
x0x00
xxx00
`

func TestSolve1(t *testing.T) {
	allTests := [][2]any{
		{INPUT_1, 140},
		{INPUT_2, 772},
		{INPUT_3, 1930},
		{INPUT_4, 5 * 12 + 10 * 18},
	}

	tutils.AssertMany(t, allTests, day12.Solve1)
}

const INPUT_5 = `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
`

const INPUT_6 = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`

func TestSolve2(t *testing.T) {
	allTests := [][2]any{
		{INPUT_1, 80},
		{INPUT_5, 236},
		{INPUT_6, 368},
		{INPUT_3, 1206},
	}

	tutils.AssertMany(t, allTests, day12.Solve2)
}
