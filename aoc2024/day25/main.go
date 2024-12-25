package day25

import (
	"strings"
)

// Returns keys, locks, then max height
func parseInput(inp string) ([][]int, [][]int, int) {
	allSch := strings.Split(inp, "\n\n")

	h := strings.Count(allSch[0], "\n") + 1
	perLine := strings.Index(allSch[0], "\n") + 1

	keys := [][]int{}
	locks := [][]int{}
	for _, v := range strings.Split(inp, "\n\n") {
		heights := make([]int, perLine)
		ct := 0

		for x := 0; x < perLine - 1; x++ {
			if v[x] == '#' {
				ct++
			}

			c := 0

			for y := 0; y < h; y++ {
				if v[y * perLine + x] == '#' {
					c++
				}
			}

			heights[x] = c
		}

		if ct == (perLine - 1) {
			keys = append(keys, heights)
		} else {
			locks = append(locks, heights)
		}
	}

	return keys, locks, h
}

func Solve1(inp string) any {
	keys, locks, h := parseInput(inp)

	c := 0
	for _, k := range keys {
		for _, l := range locks {
			good := true

			for i, h1 := range l {
				if h1 + k[i] > h {
					good = false
					break
				}
			}

			if good {
				c++
			}
		}
	}

	return c
}

func Solve2(inp string) any {
	return nil
}
