package day1

import (
	"strconv"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) []int {
	raw := strings.Split(inp, "\n")
	numbers := make([]int, 0, len(raw))

	for _, l := range raw {
		if l == "" {
			continue
		}

		v, err := strconv.Atoi(l)
		utils.PanicIfErr(err, "parsing l '%v'", l)

		numbers = append(numbers, v)
	}

	return numbers
}

func Solve1(inp string) any {
	numbers := parseInput(inp)

	for i, n1 := range numbers {
		for _, n2 := range numbers[i+1:] {
			if n1+n2 == 2020 {
				return n1 * n2
			}
		}
	}

	return nil
}

func Solve2(inp string) any {
	numbers := parseInput(inp)

	// recursion is for nerds
	for i, n1 := range numbers {
		for j, n2 := range numbers[i:] {
			if n1+n2 > 2020 {
				continue
			}

			for _, n3 := range numbers[i+j:] {
				if n1+n2+n3 == 2020 {
					return n1 * n2 * n3
				}
			}
		}
	}

	return nil
}
