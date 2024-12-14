package day14

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Robot struct {
	InitPos utils.XY
	Vel utils.XY
}

func (r Robot) AfterSeconds(s int, size utils.XY) utils.XY {
	c := r.InitPos.Add(r.Vel.Mul(s))
	c[0] %= size[0]
	c[1] %= size[1]

	if c[0] < 0 {
		c[0] += size[0]
	}

	if c[1] < 0 {
		c[1] += size[1]
	}

	return c
}

func parseInput(inp string) []*Robot {
	allRobots := []*Robot{}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l, " ")

		allRobots = append(allRobots, &Robot{
			InitPos: utils.XYFromArr(utils.SplitAndParseInt(spl[0][2:], ",")),
			Vel:     utils.XYFromArr(utils.SplitAndParseInt(spl[1][2:], ",")),
		})
	}

	return allRobots
}

func GenericSolve1(inp string, seconds, mx, my int) any {
	robots := parseInput(inp)

	size := utils.XY{mx, my}
	mid := utils.XY{mx/2, my/2}
	quadrants := [4]int{}

	for _, r := range robots {
		coords := r.AfterSeconds(seconds, size)

		if coords[0] == mid[0] || coords[1] == mid[1] {
			continue
		}

		q := 1

		if coords[0] > mid[0] {
			q++
		}

		if coords[1] > mid[1] {
			q += 2
		}

		quadrants[q - 1]++
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func Solve1(inp string) any {
	return GenericSolve1(inp, 100, 101, 103)
}

func Solve2(inp string) any {
	return nil
}
