package day10_test

import (
	"strconv"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day10"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
`

const INPUT_2 = `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
`

const INPUT_3 = `
10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
`

const INPUT_4 = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestSolve1(t *testing.T) {
	v := [][2]any{
		{INPUT_1, 2},
		{INPUT_2, 4},
		{INPUT_3, 3},
		{INPUT_4, 36},
	}

	for i, cfg := range v {
		t.Run(strconv.Itoa(i + 1), func(t *testing.T) {
			tutils.AssertFunc(t, cfg[0].(string), day10.Solve1, cfg[1])
		})
	}
}