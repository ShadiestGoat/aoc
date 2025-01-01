package day1_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2023/day1"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

const INPUT_2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT_1, day1.Solve1, 142)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT_2, day1.Solve2, 281)
}
