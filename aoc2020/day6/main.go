package day6

import (
	"strings"
)

type Group struct {
	Data map[rune]int
	Participants int
}

func parseInput(inp string) []*Group {
	allLines := strings.Split(inp, "\n\n")
	resp := []*Group{}

	for _, l := range allLines {
		data := map[rune]int{}

		count := strings.Count(l, "\n")
		l = strings.ReplaceAll(l, "\n", "")

		for _, r := range l {
			data[r]++
		}

		resp = append(resp, &Group{
			Data:         data,
			Participants: count + 1,
		})
	}

	return resp
}

func Solve1(inp string) any {
	d := parseInput(inp)
	tot := 0

	for _, g := range d {
		tot += len(g.Data)
	}

	return tot
}

func Solve2(inp string) any {
	d := parseInput(inp)
	tot := 0

	for _, g := range d {
		for _, count := range g.Data {
			if count != g.Participants {
				continue
			}

			tot++
		}
	}

	return tot
}
