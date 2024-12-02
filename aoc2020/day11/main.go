package day11

import (
	"strings"
)

const (
	STATE_FLOOR = '.'
	STATE_OCCUPIED = '#'
	STATE_EMPTY = 'L'
)

type State struct {
	LastState string
	CurState  string

	perRow int
	rowCount int
}

var SURROUND_DIRECTIONS = [][2]int{
	{0,  -1},
	{1,  -1},
	{1,   0},
	{1,   1},
	{0,   1},
	{-1,  1},
	{-1,  0},
	{-1, -1},
}

func (s State) Coord(x, y int) int {
	return y * (s.perRow + 1) + x
}

func (gs *State) Exec(maxVisDir int, tolerance int) {
	gs.LastState = gs.CurState

	newState := make([]rune, len(gs.CurState))

	s := []rune(gs.CurState)

	for i, r := range s {
		if r == STATE_FLOOR || r == '\n' {
			newState[i] = r
			continue
		}

		x := i % (gs.perRow + 1)
		y := (i - x)/(gs.perRow + 1)

		maxOccupiedAmount := tolerance

		if r == STATE_EMPTY {
			maxOccupiedAmount = 1
		}

		occupiedAmount := 0

		for _, xy := range SURROUND_DIRECTIONS {
			var finItem rune

			for visAmt := 1; visAmt <= maxVisDir; visAmt++ {
				x, y := x + xy[0] * visAmt, y + xy[1] * visAmt

				if y < 0 || x < 0 || x >= gs.perRow || y >= gs.rowCount {
					break
				}

				j := gs.Coord(x, y)
				if s[j] != STATE_FLOOR {
					finItem = s[j]
					break
				}
			}


			if finItem == STATE_OCCUPIED {
				occupiedAmount++

				if occupiedAmount >= maxOccupiedAmount {
					break
				}
			}
		}

		if r == STATE_EMPTY {
			if occupiedAmount == 0 {
				newState[i] = STATE_OCCUPIED
				continue
			}
		} else {
			if occupiedAmount >= maxOccupiedAmount {
				newState[i] = STATE_EMPTY
				continue
			}
		}

		newState[i] = r
	}

	gs.CurState = string(newState)
}

func ParseInput(inp string) *State {
	perRow := strings.Index(inp, "\n")

	return &State{
		CurState:  inp,
		perRow:    perRow,
		rowCount:  strings.Count(inp, "\n") + 1,
	}
}

func Solve1(inp string) any {
	s := ParseInput(inp)

	for {
		s.Exec(1, 4)

		if s.CurState == s.LastState {
			return strings.Count(s.CurState, string(STATE_OCCUPIED))
		}
	}
}

func Solve2(inp string) any {
	s := ParseInput(inp)
	
	// An amount thats def enough
	visAmt := s.perRow * s.rowCount

	for {
		s.Exec(visAmt, 5)

		if s.CurState == s.LastState {
			return strings.Count(s.CurState, string(STATE_OCCUPIED))
		}
	}
}
