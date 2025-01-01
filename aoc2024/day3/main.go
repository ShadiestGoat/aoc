package day3

import (
	"regexp"

	"github.com/shadiestgoat/aoc/utils/sparse"
)

// Yeah yeah yeah whatever regex is simple & I'm tired and dont want to do non-regex solution
const (
	REG_EXPR_MULT       = `mul\((-?\d{1,3}),(-?\d{1,3})\)`
	REG_EXPR_WITH_LOGIC = `(?:` + REG_EXPR_MULT + `)|(?:do\(\))|(?:don't\(\))`
)

var (
	reg_mult  = regexp.MustCompile(REG_EXPR_MULT)
	reg_logic = regexp.MustCompile(REG_EXPR_WITH_LOGIC)
)

func mult(n1, n2 string) int {
	return sparse.ParseInt(n1) * sparse.ParseInt(n2)
}

func Solve1(inp string) any {
	finds := reg_mult.FindAllStringSubmatch(inp, -1)

	s := 0
	for _, f := range finds {
		s += mult(f[1], f[2])
	}

	return s
}

func Solve2(inp string) any {
	finds := reg_logic.FindAllStringSubmatch(inp, -1)

	isEnabled := true
	s := 0

	for _, f := range finds {
		// Submatches for command would've been good but annoying to construct
		switch f[0][:3] {
		case "do(":
			isEnabled = true
		case "don":
			isEnabled = false
		case "mul":
			if isEnabled {
				s += mult(f[1], f[2])
			}
		}
	}

	return s
}
