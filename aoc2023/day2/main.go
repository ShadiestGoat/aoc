package day2

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/funiter"
	"github.com/shadiestgoat/aoc/utils/sparse"
)

var GOOD_MAP = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func parseSet(s string) (int, string) {
	spl := strings.SplitN(s, " ", 2)
	return sparse.ParseInt(spl[0]), spl[1]
}

var (
	scanSet = funiter.NewScanFunc(", ")
	scanGame = funiter.NewScanFunc("; ")
	finAdder = funiter.NewScanSumFunc("\n")
)

func addFor1(game string) int {
	g := strings.SplitN(game, ": ", 2)

	bad := scanGame(g[1], func(set string) bool {
		return scanSet(set, func (s string) bool {
			n, c := parseSet(s)

			return n > GOOD_MAP[c]
		})
	})

	if bad {
		return 0
	}

	return sparse.ParseInt(g[0][5:])
}

func addFor2(game string) int {
	m := map[string]int{}
	t := 0

	scanGame(strings.SplitN(game, ": ", 2)[1], func (set string) bool {
		return scanSet(set, func (v string) bool {
			n, c := parseSet(v)
			if m[c] == 0 {
				t += n				
			} else if n > m[c] {
				t += n - m[c]
				m[c] = n
			}

			return false
		})
	})

	return t
}

func Solve1(inp string) any {
	return finAdder(inp, addFor1)
}

func Solve2(inp string) any {
	return finAdder(inp, addFor2)
}
