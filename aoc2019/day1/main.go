package day1

import (
	"strconv"
	"strings"

	"shadygoat.eu/aoc/utils"
)

func parseLine(l string) int {
	if l == "" {
		return 0
	}

	v, err := strconv.Atoi(l)
	utils.PanicIfErr(err, "parsing line '%v'", l)

	return v/3 - 2
}

func Solve1(inp string) any {
	t := 0

	for _, l := range strings.Split(inp, "\n") {
		t += parseLine(l)
	}

	return t
}

func Solve2(inp string) any {
	return nil
}
