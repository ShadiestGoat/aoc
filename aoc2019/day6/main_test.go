package day6_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day6"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`

const INPUT_2 = `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day6.Solve1, 42)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT_2, day6.Solve2, 4)
}
