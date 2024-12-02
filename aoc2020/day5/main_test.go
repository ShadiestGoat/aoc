package day5_test

import (
	"testing"

	"shadygoat.eu/aoc/aoc2020/day5"
)

func TestParseRow(t *testing.T) {
	testData := []string{
		`FBFBBFFRLR`,
		`BFFFBBFRRR`,
		`FFFBBBFRRR`,
		`BBFFBBFRLL`,
	}
	expected_results := [][2]int{
		{44, 5},
		{70, 7},
		{14, 7},
		{102, 4},
	}

	for i, d := range testData {
		resp := day5.ParseRow(d)
		exp := expected_results[i]

		if resp[0] != exp[0] || resp[1] != exp[1] {
			t.Fatalf("Failed to parse '%v': Expected %#v, got %#v", d, exp, resp)
		}
	}
}
