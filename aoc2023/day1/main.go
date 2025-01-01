package day1

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/funiter"
)

var nums = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func parseN(str string, includeFakes bool) int {
	r := str[0]
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}

	if !includeFakes {
		return -1
	}

	for i, n := range nums {
		if strings.HasPrefix(str, n) {
			return i + 1
		}
	}

	return -1
}

func genericSolve(inp string, includeFakes bool) int {
	c := 0

	funiter.SplitAndScan(inp, "\n", func(s string) {
		v := 0

		for i := 0; i < len(s); i++ {
			if n := parseN(s[i:], includeFakes); n != -1 {
				v = n * 10

				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if n := parseN(s[i:], includeFakes); n != -1 {
				v += n

				break
			}
		}

		c += v
	})

	return c
}

func Solve1(inp string) any {
	return genericSolve(inp, false)
}

func Solve2(inp string) any {
	return genericSolve(inp, true)
}
