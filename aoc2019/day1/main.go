package day1

import (
	"github.com/shadiestgoat/aoc/utils/sparse"
)

func parseInput(inp string) []int {
	return sparse.SplitAndParseInt(inp, "\n")
}

func fuelRequirements(fuel int) int {
	return fuel/3 - 2
}

func totFuelRequirements(fuel int) int {
	last := fuelRequirements(fuel)
	t := 0

	for last > 0 {
		t += last

		last = fuelRequirements(last)
	}

	return t
}

func Solve1(inp string) any {
	allFuel := parseInput(inp)
	t := 0

	for _, f := range allFuel {
		t += fuelRequirements(f)
	}

	return t
}

func Solve2(inp string) any {
	allFuel := parseInput(inp)
	t := 0

	for _, f := range allFuel {
		t += totFuelRequirements(f)
	}

	return t
}
