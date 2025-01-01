package day3

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/xy"
)

var (
	DIR_TRANSLATION = map[rune]xy.XY{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, -1},
		'D': {0, 1},
	}
)

type Instruction struct {
	Amount int
	Dir    xy.XY
}

func parseInput(inp string) [2][]*Instruction {
	spl := strings.Split(inp, "\n")
	o := [2][]*Instruction{}

	for i, l := range spl {
		for _, ins := range strings.Split(l, ",") {
			dir := DIR_TRANSLATION[rune(ins[0])]
			o[i] = append(o[i], &Instruction{
				Amount: sparse.ParseInt(ins[1:]),
				Dir:    dir,
			})
		}
	}

	return o
}

func (ins *Instruction) DoSteps(cur xy.XY, h func(pos xy.XY)) xy.XY {
	for i := 1; i <= ins.Amount; i++ {
		cur = cur.Add(ins.Dir)
		h(cur)
	}

	return cur
}

func doInsSet(ins []*Instruction, h func(pos xy.XY)) {
	curPos := xy.XY{}

	for _, v := range ins {
		curPos = v.DoSteps(curPos, h)
	}
}

func Solve1(inp string) any {
	ins := parseInput(inp)
	wireAPath := map[xy.XY]bool{}

	doInsSet(ins[0], func(pos xy.XY) {
		wireAPath[pos] = true
	})

	delete(wireAPath, xy.XY{})

	intersection := xy.XY{}

	doInsSet(ins[1], func(pos xy.XY) {
		if !wireAPath[pos] {
			return
		}

		if intersection.IsAtOrigin() || pos.ManhattanDistance() < intersection.ManhattanDistance() {
			intersection = pos
		}
	})

	return intersection.ManhattanDistance()
}

func Solve2(inp string) any {
	ins := parseInput(inp)
	wireAPath := map[xy.XY]int{}
	steps := 0

	doInsSet(ins[0], func(pos xy.XY) {
		steps++

		if wireAPath[pos] == 0 {
			wireAPath[pos] = steps
		}
	})

	steps = 0
	bestStepCount := 0
	doInsSet(ins[1], func(pos xy.XY) {
		steps++

		if wireAPath[pos] == 0 {
			return
		}

		intStepCount := steps + wireAPath[pos]
		if bestStepCount == 0 || intStepCount < bestStepCount {
			bestStepCount = intStepCount
		}
	})

	return bestStepCount
}
