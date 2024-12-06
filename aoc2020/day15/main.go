package day15

import (
	"strconv"
	"strings"
)

func parseInput(inp string) []int {
	spl := strings.Split(inp, ",")
	nums := make([]int, len(spl))

	for i, v := range spl {
		n, _ := strconv.Atoi(v)
		nums[i] = n
	}

	return nums
}

func PlayGame(inp string, max int) int {
	starting := parseInput(inp)

	his := map[int][2]int{}

	gi := 1

	addToHis := func(v int) {
		old := his[v]

		old[0], old[1] = gi, old[0]
		his[v] = old
	}

	for _, v := range starting {
		addToHis(v)
		gi++
	}

	lastNum := starting[len(starting)-1]

	for {
		if gi == max+1 {
			return lastNum
		}

		spoken := 0

		numHis := his[lastNum]
		if numHis[1] != 0 {
			spoken = numHis[0] - numHis[1]
		}

		addToHis(spoken)
		lastNum = spoken

		gi++
	}
}

func Solve1(inp string) any {
	return PlayGame(inp, 2020)
}

func Solve2(inp string) any {
	return PlayGame(inp, 30000000)
}
