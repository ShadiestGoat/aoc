package day3_test

import (
	"testing"

	"shadygoat.eu/aoc/aoc2020/day3"
)

const INPUT = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestSolve1(t *testing.T) {
	v := day3.Solve1(INPUT)

	if v != 7 {
		t.Fatalf("Bad input - expected 7, got %v", v)
	}
}
