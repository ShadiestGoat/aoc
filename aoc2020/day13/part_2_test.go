package day13_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day13"
)

func TestGenericCRA(t *testing.T) {
	// geek4geek example <3
	num := []int{3, 4, 5}
	rem := []int{2, 3, 1}

	resp := day13.GenericCRA(num, rem, day13.Product(num))

	if resp != 11 {
		t.Fatalf("What the hecc (11 (e) != %v (r))", resp)
	}
}

func TestSolve2(t *testing.T) {
	inp := []string{
		`17,x,13,19`,
		`67,7,59,61`,
		`67,x,7,59,61`,
		`67,7,x,59,61`,
		`1789,37,47,1889`,
	}
	exp := []int{
		3417,
		754018,
		779210,
		1261476,
		1202161486,
	}

	for i, v := range inp {
		resp := day13.Solve2("\n" + v)

		if resp != exp[i] {
			t.Errorf("Oof... expected %v, got %v", exp[i], resp)
		}
	}
}
