package day9_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day9"
	"github.com/shadiestgoat/aoc/utils/sparse"
)

const INPUT = `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

func TestGenericSolver1(t *testing.T) {
	resp := day9.GenericSolver1(strings.TrimSpace(INPUT), 5)

	if resp != 127 {
		t.Fatalf("Grr >:( expected 127, got %v", resp)
	}
}

func TestGenericSolver2(t *testing.T) {
	min, max := day9.GenericSolver2(sparse.SplitAndParseInt(strings.TrimSpace(INPUT), "\n"), 127)

	if min != 2 || max != 5 {
		t.Fatalf("Grr >:( expected (2, 5), got %v", [2]int{min, max})
	}
}
