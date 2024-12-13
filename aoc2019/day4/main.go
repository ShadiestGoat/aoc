package day4

import (
	"strconv"

	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) []int {
	return utils.SplitAndParseInt(inp, "-")
}

func verifyNum(n int, noBig bool) bool {
	doubleFound := false

	str := strconv.Itoa(n)
	for i, r := range str[1:] {
		l := rune(str[i])
		if r < l {
			return false
		}

		if l == r {
			if noBig {
				if i > 0 && rune(str[i - 1]) == r {
					continue
				}
				if i < len(str) - 2 && rune(str[i + 2]) == r {
					continue
				}
			}

			doubleFound = true
		}
	}

	return doubleFound
}

func testNumbers(min, max int, noBig bool) int {
	c := 0

	for i := min; i <= max; i++ {
		if verifyNum(i, noBig) {
			c++
		}
	}

	return c
}

func Solve1(inp string) any {
	inpRange := parseInput(inp)

	return testNumbers(inpRange[0], inpRange[1], false)
}

func Solve2(inp string) any {
	inpRange := parseInput(inp)

	return testNumbers(inpRange[0], inpRange[1], true)
}
