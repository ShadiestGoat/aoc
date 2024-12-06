package day6

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type State struct {
	// Coord -> Directions
	DirHist map[utils.XY][]utils.XY
	CurPos  utils.XY
	Dir     utils.XY
	Board   []string
}

func (s State) outOfBounds(c utils.XY) bool {
	return c.OutOfBounds(utils.XY{len(s.Board[0]) - 1, len(s.Board) - 1})
}

// Does 1 action. Either moves or rotates
// Returns true if goes of the map
// Doesn't record history
func (s *State) GoOnce() bool {
	nextV := s.CurPos.Add(s.Dir)

	if s.outOfBounds(nextV) {
		return true
	}

	if s.Board[nextV[1]][nextV[0]] == '#' {
		s.Dir = s.Dir.RotateUnitVector(2)

		return false
	}

	s.CurPos = nextV

	return false
}

func (s *State) GoUntilExit() {
	for !s.GoOnce() {
		s.DirHist[s.CurPos] = append(s.DirHist[s.CurPos], s.Dir)
	}
}

func (s *State) findObstacleInDirection(og utils.XY, dir utils.XY) bool {
	i := 1
	for {
		pos := og.Add(dir.Mul(i))

		if s.outOfBounds(pos) {
			return false
		}

		if s.Board[pos[1]][pos[0]] == '#' {
			return true
		}

		i++
	}
}

func parseInput(inp string) *State {
	rawPos := strings.Index(inp, "^")
	// Its been a tiring week... I'm not gonna bother with a fully string board like aoc2020/day11
	lines := strings.Split(strings.Replace(inp, "^", ".", 1), "\n")
	perLineChars := len(lines[0]) + 1

	pos := utils.XY{rawPos % perLineChars, rawPos / perLineChars}

	return &State{
		DirHist: map[utils.XY][]utils.XY{
			pos: {
				{0, -1},
			},
		},
		CurPos: pos,
		Dir:    utils.XY{0, -1},
		Board:  lines,
	}
}

func Solve1(inp string) any {
	s := parseInput(inp)
	s.GoUntilExit()

	return len(s.DirHist)
}

func Solve2(inp string) any {
	s := parseInput(inp)

	ogPos := s.CurPos
	ogBoard := make([]string, len(s.Board))
	copy(ogBoard, s.Board)

	s.GoUntilExit()

	oldHist := s.DirHist
	delete(oldHist, ogPos)
	testHist := map[utils.XY]bool{}
	dirs := []utils.XY{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for c := range oldHist {
		for _, d := range dirs {
			if s.findObstacleInDirection(c, d) {
				testHist[c] = true
				break
			}
		}
	}

	amt := 0

	for c := range testHist {
		// Reset to og
		s.DirHist = map[utils.XY][]utils.XY{
			ogPos: {
				{0, -1},
			},
		}

		s.Dir = utils.XY{0, -1}
		s.CurPos = ogPos
		s.Board = make([]string, len(s.Board))
		copy(s.Board, ogBoard)

		row := s.Board[c[1]]
		s.Board[c[1]] = row[:c[0]] + "#" + row[c[0]+1:]

		for !s.GoOnce() {
			if slices.Contains(s.DirHist[s.CurPos], s.Dir) {
				amt++
				break
			}

			s.DirHist[s.CurPos] = append(s.DirHist[s.CurPos], s.Dir)
		}
	}

	return amt
}
