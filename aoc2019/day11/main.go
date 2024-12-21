package day11

import (
	"slices"

	"github.com/shadiestgoat/aoc/aoc2019/intcode"
	"github.com/shadiestgoat/aoc/utils"
)

type Robot struct {
	Dir utils.XY
	Pos utils.XY
	comp *intcode.Computer
}

func parseInput(inp string) *Robot {
	return &Robot{
		Dir:  utils.DIR_UP,
		Pos:  utils.XY{},
		comp: &intcode.Computer{
			Code:   intcode.ParseIntCode(inp),
		},
	}
}

// Return true if the program completes
func (r *Robot) RunOnce(w map[utils.XY]int) bool {
	r.comp.Input = append(r.comp.Input, w[r.Pos])
	exitEarly := r.comp.RunIntCode()

	w[r.Pos] = r.comp.Output[0]

	rot := -2
	if r.comp.Output[1] == 1 {
		rot = 2
	}

	r.comp.Output = r.comp.Output[2:]

	r.Dir = r.Dir.RotateUnitVector(rot)
	r.Pos = r.Pos.Add(r.Dir)

	return exitEarly
}

func Solve1(inp string) any {
	r := parseInput(inp)
	w := map[utils.XY]int{}
	for !r.RunOnce(w) {

	}

	return len(w)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func Solve2(inp string) any {
	r := parseInput(inp)
	w := map[utils.XY]int{
		r.Pos: 1,
	}
	for !r.RunOnce(w) {

	}

	min := utils.XY{}
	max := utils.XY{}
	inited := false

	for p := range w {
		if !inited || p[0] < min[0] {
			min[0] = p[0]
		}
		if !inited || p[1] < min[1] {
			min[1] = p[1]
		}
		if !inited || p[0] > max[0] {
			max[0] = p[0]
		}
		if !inited || p[1] > max[1] {
			max[1] = p[1]
		}

		inited = true
	}

	sizeX := diff(min[0], max[0]) + 1
	sizeY := diff(min[1], max[1]) + 1

	str := make([][]rune, sizeY)
	row := make([]rune, sizeX)
	for i := 0; i < sizeX; i++ {
		row[i] = ' '
	}
	for i := range str {
		str[i] = slices.Clone(row)
	}

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			c := utils.XY{x, y}.Add(min.Mul(-1))
			if w[c] == 0 {
				continue
			}

			str[y][x] = '█'
		}
	}

	o := ""

	for _, r := range str {
		o += "\n" + string(r)
	}

	return o[1:]
}
