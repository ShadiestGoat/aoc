package day14

import (
	"fmt"
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils/mutils"
	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/xy"
)

type Robot struct {
	Pos xy.XY
	Vel xy.XY
}

func (r Robot) AfterSeconds(s int, size xy.XY) xy.XY {
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

func parseInput(inp string) []*Robot {
	allRobots := []*Robot{}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l, " ")

		allRobots = append(allRobots, &Robot{
			Pos: xy.XYFromArr(sparse.SplitAndParseInt(spl[0][2:], ",")),
			Vel: xy.XYFromArr(sparse.SplitAndParseInt(spl[1][2:], ",")),
		})
	}

	return allRobots
}

func moveAllRobots(allRobots []*Robot, mx, my int) {
	for _, r := range allRobots {
		r.Pos = r.AfterSeconds(1, xy.XY{mx, my})
	}
}

func GenericSolve1(inp string, seconds, mx, my int) any {
	robots := parseInput(inp)

	size := xy.XY{mx, my}
	mid := xy.XY{mx / 2, my / 2}
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

		quadrants[q-1]++
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

// Trust me, thats what were doing
func differentiate(cur []int) []int {
	diff := make([]int, len(cur) - 1)

	for j := 0; j < len(cur) - 1; j++ {
		diff[j] = cur[j + 1] - cur[j]
	}

	return diff
}

func clamp(min, v, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

func Solve2(inp string) any {
	robots := parseInput(inp)

	i := 1.0
	totMH := 0.0
	interestingPos := []int{}
	for {
		moveAllRobots(robots, 101, 103)
		curMH := 0.0

		for i, r := range robots {
			for _, r2 := range robots[i + 1:] {
				curMH += float64(r2.Pos.ManhattanDistanceTo(r.Pos))
			}
		}

		if i > 100 && curMH < (totMH/i)*0.85 {
			interestingPos = append(interestingPos, int(i))

			if len(interestingPos) == 5 {
				break
			}
		}
		if i > 100 && curMH < (totMH/i)*0.65 {
			return i
		}

		totMH += curMH
		i++
	}

	diff1 := differentiate(interestingPos)

	// The current number
	cur := interestingPos[3]
	// The current first order diff
	curDiff1 := diff1[2]

	// The current second order diffs
	diff2 := differentiate(diff1)
	if diff2[0] < 0 {
		diff2 = diff2[1:]
		cur = interestingPos[4]
		curDiff1 = diff1[3]
	} else {
		diff2 = diff2[:len(diff2) - 1]
	}

	diff2Diff := [2]int{
		mutils.Dir(diff2[0]) * 4,
		mutils.Dir(diff2[1]) * 4,
	}

	found := 0
	for {
		didQuirky := false

		for i := 0; i < 2; i++ {
			diff2[i] += diff2Diff[i]

			if diff2[i] <= -100 || diff2[i] >= 100 {
				didQuirky = true
				diff2Diff[i] = -diff2Diff[i]

				found++
				if found > 2 {
					return cur
				}
			}

			curDiff1 += clamp(-100, diff2[i], 100)
			cur += curDiff1
		}

		if didQuirky {
			diff2[0], diff2[1] = -diff2[1], -diff2[0]
		}
	}
}
