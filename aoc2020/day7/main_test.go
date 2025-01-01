package day7_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day7"
	"github.com/shadiestgoat/aoc/utils/xprint"
)

const INPUT = `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

func TestSolve1(t *testing.T) {
	d := day7.Solve1(strings.TrimSpace(INPUT))

	if d != 4 {
		t.Fatalf(":( Expected 4, got %v", d)
	}
}

const INPUT_2 = `
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

func TestSolve2(t *testing.T) {
	inp := day7.ParseInput(strings.TrimSpace(INPUT_2))
	xprint.PrintJSON(inp)

	d := day7.Solve2(strings.TrimSpace(INPUT_2))

	if d != 126 {
		t.Fatalf("Ah :( Expected 126, got %v", d)
	}
}
