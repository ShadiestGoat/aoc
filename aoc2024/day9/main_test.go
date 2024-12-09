package day9_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day9"
	"github.com/shadiestgoat/aoc/utils/tutils"
)


func TestSolve1(t *testing.T) {
    tutils.AssertFunc(t, `2333133121414131402`, day9.Solve1, 1928)
}

func TestSolve2(t *testing.T) {
    tutils.AssertFunc(t, `2333133121414131402`, day9.Solve2, 2858)
}
