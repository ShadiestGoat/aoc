package day13

import (
	"fmt"

	"github.com/shadiestgoat/aoc/aoc2019/intcode"
	"github.com/shadiestgoat/aoc/utils/xy"
)

func initialMap(inp string) map[xy.XY]int {
	o := intcode.QuickRun(intcode.ParseIntCode(inp), nil)
	return parseOutput(o)
}

func parseOutput(o []int) map[xy.XY]int {
	obj := map[xy.XY]int{}
	for i := 0; i < len(o); i += 3 {
		obj[xy.XY{o[i], o[i+1]}] = o[i+2]
	}

	return obj
}

func Solve1(inp string) any {
	c := 0
	for _, t := range initialMap(inp) {
		if t == 2 {
			c++
		}
	}

	return c
}

var tileMapTrans = [5]rune{' ', '█', '▣', '▔', 'O'}

func drawMap(base map[xy.XY]int, player, ball xy.XY) {
	maxX, maxY := 0, 0

	for c := range base {
		if c[0] > maxX {
			maxX = c[0]
		}

		if c[1] > maxY {
			maxY = c[1]
		}
	}

	lines := ""
	for y := 0; y <= maxY; y++ {
		l := make([]rune, maxX+1)
		for x := 0; x <= maxX; x++ {
			l[x] = tileMapTrans[base[xy.XY{x, y}]]
		}

		if player[1] == y {
			l[player[0]] = tileMapTrans[3]
		}
		if ball[1] == y {
			l[ball[0]] = tileMapTrans[4]
		}

		lines += "\n" + string(l)
	}

	fmt.Println(lines[1:])
}

func Solve2(inp string) any {
	code := intcode.ParseIntCode(inp)
	code[0] = 2

	comp := &intcode.Computer{
		Input: []int{},
		Code:  code,
	}
	comp.RunIntCode()

	playerPos := xy.XY{}
	ballPos := xy.XY{}
	for i := 0; i < len(comp.Output); i += 3 {
		if comp.Output[i+2] == 3 {
			playerPos = xy.XY{comp.Output[i], comp.Output[i+1]}
		} else if comp.Output[i+2] == 4 {
			ballPos = xy.XY{comp.Output[i], comp.Output[i+1]}
		}

		if !playerPos.IsAtOrigin() && !ballPos.IsAtOrigin() {
			break
		}
	}

	m := parseOutput(comp.Output)
	m[playerPos] = 0
	m[ballPos] = 0

	comp.Output = []int{-1, -1, 0, ballPos[0], ballPos[1], 4}
	comp.Input = []int{}
	score := 0

	for {
		for i := 0; i < len(comp.Output); i += 3 {
			if comp.Output[i+2] == 0 {
				delete(m, xy.XY{comp.Output[i], comp.Output[i+1]})
			} else if comp.Output[i] == -1 && comp.Output[i+1] == 0 {
				score = comp.Output[i+2]
			}
		}

		hasBlock := false
		for _, v := range m {
			if v == 2 {
				hasBlock = true
				break
			}
		}

		if !hasBlock {
			break
		}

		// lastPos := xy.XY{comp.Output[0], comp.Output[1]}
		ballPos := xy.XY{comp.Output[len(comp.Output)-3], comp.Output[len(comp.Output)-2]}
		diff := ballPos.Add(playerPos.Mul(-1)).Unit()

		comp.Input = append(comp.Input, diff[0])
		comp.Output = []int{}
		playerPos[0] += diff[0]
		comp.RunIntCode()
	}

	return score
}
