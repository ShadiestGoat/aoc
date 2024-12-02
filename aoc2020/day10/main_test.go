package day10_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day10"
)

const INPUT_1 = `
16
10
15
5
1
11
7
19
6
12
4
`

const INPUT_2 = `
1
2
3
4
7
8
9
10
11
14
17
18
19
20
23
24
25
28
31
33
32
34
35
38
39
42
45
46
47
48
49
`

func TestSolver2(t *testing.T) {
	inputs := []string{INPUT_1, INPUT_2}
	exp := []int{8, 19208}

	for i, inp := range inputs {
		resp := day10.Solve2(strings.TrimSpace(inp))

		if resp != exp[i] {
			t.Errorf("Failed in example #%v: Expected %v, got %v", i+1, exp[i], resp)
		}
	}
}
