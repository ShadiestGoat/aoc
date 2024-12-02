package day2_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day2"
	"github.com/shadiestgoat/aoc/utils"
)

const CASES = `
7 6 4 2 1: S
1 2 7 8 9: U
9 7 6 2 1: U
1 3 2 4 5: S
8 6 4 4 1: S
1 3 6 7 9: S
`

func TestIsSafeWithoutOne(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(CASES), "\n")

	for _, l := range lines {
		t.Run(l, func(t *testing.T) {
			spl := strings.Split(l, ": ")

			r := utils.SplitAndParseInt(spl[0], " ")

			e := spl[1] == "S"
			safe := day2.IsSafeWithoutOne(r)

			if e != safe {
				t.Errorf("Fail: Expected safe: %t; got %t", e, safe)
			}
		})
	}
}
