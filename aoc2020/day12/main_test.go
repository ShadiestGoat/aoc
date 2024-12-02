package day12_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day12"
)

const INPUT = `
F10
N3
F7
R90
F11
`

func TestSolve1(t *testing.T) {
	resp := day12.Solve1(strings.TrimSpace(INPUT))

	if resp != 25 {
		t.Fatalf("Grr :( -- expected 25, got %v", resp)
	}
}

func TestSolve2(t *testing.T) {
	resp := day12.Solve2(strings.TrimSpace(INPUT))

	if resp != 286 {
		t.Fatalf("Mrmrm expected 286, got %v", resp)
	}
}
