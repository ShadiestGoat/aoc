package day17_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day17"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
Register A: 10

Program: 5,0,5,1,5,4
`

const INPUT_2 = `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

const INPUT_3 = `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`

const INPUT_4 = `
Register A: 117440
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`

func TestSolve1(t *testing.T) {
	allTests := [][2]any{
		{INPUT_1, `0,1,2`},
		{INPUT_2, `4,6,3,5,6,3,5,2,1,0`},
		{INPUT_4, `0,3,5,4,3,0`},
	}

	tutils.AssertMany(t, allTests, day17.Solve1)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT_3, day17.Solve2, 117440)
}
