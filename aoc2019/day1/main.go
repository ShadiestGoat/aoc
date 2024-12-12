package day1

import (
	"strconv"

	"github.com/shadiestgoat/aoc/utils"
)

func parseLine(l string) int {
	if l == "" {
		return 0
	}

	v, err := strconv.Atoi(l)
	utils.PanicIfErr(err, "parsing line '%v'", l)

	return v/3 - 2
}

func parseInput(inp string) []int {
	return utils.SplitAndParseInt(inp, "\n")
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
