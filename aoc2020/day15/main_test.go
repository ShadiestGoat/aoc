package day15_test

import (
	"strconv"
	"strings"
	"testing"

	"shadygoat.eu/aoc/aoc2020/day15"
)

func TestSolve1(t *testing.T) {
	cfg := []string{
		`1,3,2:1`,
		`2,1,3:10`,
		`1,2,3:27`,
		`2,3,1:78`,
		`3,2,1:438`,
		`3,1,2:1836`,
	}

	for _, v := range cfg {
		spl := strings.Split(v, ":")
		exp, _ := strconv.Atoi(spl[1])

		resp := day15.Solve1(spl[0])

		if resp != exp {
			t.Errorf("Grr >:( failed example %v: Expected %v, got %v", spl[0], exp, resp)
		}
	}
}