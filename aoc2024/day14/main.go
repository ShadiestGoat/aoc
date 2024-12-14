package day14

import (
	"fmt"
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Robot struct {
	Pos utils.XY
	Vel utils.XY
}

func (r Robot) AfterSeconds(s int, size utils.XY) utils.XY {
	c := r.Pos.Add(r.Vel.Mul(s))
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

func (r Robot) FindNearestNeighbor(dir utils.XY) int {
	for {

	}
}

func parseInput(inp string) []*Robot {
	allRobots := []*Robot{}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l, " ")

		allRobots = append(allRobots, &Robot{
			Pos: utils.XYFromArr(utils.SplitAndParseInt(spl[0][2:], ",")),
			Vel:     utils.XYFromArr(utils.SplitAndParseInt(spl[1][2:], ",")),
		})
	}

	return allRobots
}

func moveAllRobots(allRobots []*Robot, mx, my int) {
	for _, r := range allRobots {
		r.Pos = r.AfterSeconds(1, utils.XY{mx, my})
	}
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

func drawOutput(robots []*Robot, mx, my int) {
	row := []rune(strings.Repeat(" ", mx))

	str := [][]rune{}

	for i := 0; i < my; i++ {
		str = append(str, slices.Clone(row))
	}

	for _, r := range robots {
		str[r.Pos[1]][r.Pos[0]] = '█'
	}

	o := ""

	for _, row := range str {
		o += "\n" + string(row)
	}

	fmt.Println(o[1:])
}

func Solve2(inp string) any {
	robots := parseInput(inp)

	tot := 0.0
	i := 0.0

	for {
		moveAllRobots(robots, 101, 103)
		cur := 0

		for j, r := range robots {
			for _, tr := range robots[j:] {
				cur += r.Pos.ManhattanDistanceTo(tr.Pos)
			}
		}

		// How luck based is this? Ssshhhh
		if i > 50 && float64(cur) < (tot/i)*0.6 {
			drawOutput(robots, 101, 103)
			return i + 1
		}

		tot += float64(cur)
		i++
	}
}
