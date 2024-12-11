package day24

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

var (
	ALL_DIRS = []utils.XY{
		{-2, 0},
		{2, 0},
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}
)

func isModifier(r rune) bool {
	return r == 's' || r == 'n'
}

func ParseInput(inp string) [][]utils.XY {
	lines := strings.Split(inp, "\n")
	o := [][]utils.XY{}

	for _, l := range lines {
		dirs := []utils.XY{}

		for i, r := range l {
			if isModifier(r) {
				continue
			}
	
			x, y := 1, 0
			if r == 'w' {
				x = -1
			}

			lastIsModifier := false
			if i > 0 {
				lastR := rune(l[i - 1])

				if isModifier(lastR) {
					lastIsModifier = true

					if lastR == 'n' {
						y = -1
					} else {
						y = 1
					}
				}
			}

			if !lastIsModifier {
				x *= 2
			}

			dirs = append(dirs, utils.XY{x, y})
		}

		o = append(o, dirs)
	}

	return o
}

func DoDirs(dirs []utils.XY) utils.XY {
	cur := utils.XY{}

	for _, v := range dirs {
		cur = cur.Add(v)
	}

	return cur
}

func createInitialMap(dirs [][]utils.XY) map[utils.XY]bool {
	o := map[utils.XY]bool{}

	for _, d := range dirs {
		f := DoDirs(d)

		if o[f] {
			delete(o, f)
		} else {
			o[f] = true
		}
	}

	return o
}

func Solve1(inp string) any {
	dirs := ParseInput(inp)
	m := createInitialMap(dirs)

	return len(m)
}

func countBlackTileAroundCoord(state map[utils.XY]bool, coord utils.XY) int {
	c := 0

	for _, d := range ALL_DIRS {
		if state[coord.Add(d)] {
			c++
		}
	}

	return c
}

func playGame(state map[utils.XY]bool, moves int) map[utils.XY]bool {
	for i := 0; i < moves; i++ {
		newState := map[utils.XY]bool{}

		for b, sc := range state {
			if !sc {
				continue
			}

			bCount := countBlackTileAroundCoord(state, b)
			if bCount == 1 || bCount == 2 {
				newState[b] = true
			}

			for _, d := range ALL_DIRS {
				wCoord := b.Add(d)
				if countBlackTileAroundCoord(state, wCoord) == 2 {
					newState[wCoord] = true
				}
			}
		}

		state = newState
	}

	return state
}

func Solve2(inp string) any {
	m := createInitialMap(ParseInput(inp))
	m = playGame(m, 100)

	return len(m)
}