package day2

import (
	"math"

	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) [][]int {
	return utils.SplitAndParseFunc(inp, "\n", func(s string) []int {
		return utils.SplitAndParseInt(s, " ")
	})
}

func isSafe(r []int) bool {
	increasing := r[0] < r[1]

	for li, c := range r[1:] {
		if increasing != (r[li] < c) {
			return false
		}

		d := math.Abs(float64(r[li] - c))

		if d < 1 || d > 3 {
			return false
		}
	}

	return true
}

func Solve1(inp string) any {
	safes := 0
	p := parseInput(inp)

	for _, r := range p {
		if isSafe(r) {
			safes++
		}
	}

	return safes
}

func IsSafeWithoutOne(r []int) bool {
	for i := 0; i < len(r); i++ {
		cr := make([]int, len(r)-1)
		copy(cr[:i], r[:i])
		copy(cr[i:], r[i+1:])

		if isSafe(cr) {
			return true
		}
	}

	return false
}

func Solve2(inp string) any {
	safes := 0
	p := parseInput(inp)

	for _, r := range p {
		if isSafe(r) || IsSafeWithoutOne(r) {
			safes++
			continue
		}
	}

	return safes
}
