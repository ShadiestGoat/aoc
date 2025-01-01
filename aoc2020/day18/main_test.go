package day18_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day18"
	"github.com/shadiestgoat/aoc/utils/sparse"
)

func doTest(t *testing.T, tests []string, h func(string) int) {
	for _, te := range tests {
		cfg := strings.Split(te, " = ")

		resp := h(cfg[0])

		if resp != sparse.ParseInt(cfg[1]) {
			t.Errorf("Failed to parse '%v': expected %v, got %v", cfg[0], cfg[1], resp)
		}
	}
}

func TestDoMathLTR(t *testing.T) {
	tests := []string{
		`1 + 2 = 3`,
		`1 * 2 = 2`,
		`2 * 2 = 4`,
		`1 + 2 * 3 = 9`,
		`12 + 123 * 2456 = 331560`,
		`1 + (2 * 3) = 7`,
		`(1 + 2) * 3 = 9`,
		`(1 + 2) + (1 + 2) = 6`,
	}

	doTest(t, tests, day18.DoMathLTR)
}

func TestDoMathPlus(t *testing.T) {
	tests := []string{
		// `1 + 2 * 3 + 4 = 21`,
		// `1 * 2 = 2`,
		// `1 + 2 * 3 = 9`,
		// `12 * 123 + 2456 = 30948`,
		// `1 + (2 * 3) = 7`,
		// `(1 + 2) * 3 = 9`,
		// `(1 + 2) + (1 + 2) = 6`,
		// `1 + (2 * 3) + (4 * (5 + 6)) = 51`,
		// `2 * 3 + (4 * 5) = 46`,
		`5 + (8 * 3 + 9 + 3 * 4 * 3) = 1445`,
		// `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) = 669060`,
		// `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 = 23340`,
	}

	doTest(t, tests, day18.DoMathPlus)
}
