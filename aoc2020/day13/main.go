package day13

import (
	"strconv"
	"strings"
)

func parseInput(inp string) any {
	return nil
}

func Solve1(inp string) any {
	spl := strings.Split(inp, "\n")

	departTime, _ := strconv.Atoi(spl[0])

	minTime := departTime
	minID := 0

	for _, t := range strings.Split(spl[1], ",") {
		if t == "x" {
			continue
		}

		id, _ := strconv.Atoi(t)
		waitTime := id - departTime%id

		if waitTime < minTime {
			minTime = waitTime
			minID = id
		}
	}

	return minID * minTime
}
