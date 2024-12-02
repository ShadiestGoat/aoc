package day1

import (
	"math"
	"slices"
	"strings"

	"shadygoat.eu/aoc/utils"
)

func parseInput(inp string, h func (vl, vr int)) {
	lines := strings.Split(inp, "\n")

	for _, s := range lines {
		spl := strings.Split(s, "   ")

		h(utils.ParseInt(spl[0]), utils.ParseInt(spl[1]))
	}
}

func Solve1(inp string) any {
	l, r := []int{}, []int{}
	parseInput(inp, func(vl, vr int) {
		l = append(l, vl)
		r = append(l, vr)
	})

	slices.Sort(l)
	slices.Sort(r)

	s := 0
	for i := 0; i < len(l); i++ {
		s += int(math.Abs(float64(l[i] - r[i])))
	}

	return s
}

func Solve2(inp string) any {
	l, r := []int{}, map[int]int{}
	parseInput(inp, func(vl, vr int) {
		l = append(l, vl)
		r[vr]++
	})

	s := 0
	for _, v := range l {
		s += v * r[v]
	}

	return s
}
