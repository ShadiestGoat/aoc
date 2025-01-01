package day10

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/xy"
)

var (
	ALL_DIRS = []xy.XY{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

type State struct {
	Lines []string
	size  xy.XY
}

func parseInput(inp string) *State {
	lines := strings.Split(inp, "\n")
	return &State{
		Lines: lines,
		size:  xy.GetSizeString(lines),
	}
}

func (s *State) valAt(c xy.XY) rune {
	return rune(s.Lines[c[1]][c[0]])
}

func (s *State) CountPaths(cur xy.XY, counted map[xy.XY]bool) int {
	curV := s.valAt(cur)
	if curV == '9' {
		if counted != nil {
			if counted[cur] {
				return 0
			}
			counted[cur] = true
		}

		return 1
	}

	c := 0
	for _, d := range ALL_DIRS {
		n := cur.Add(d)
		if n.OutOfBounds(s.size) {
			continue
		}

		v := s.valAt(n)

		if v == '.' {
			continue
		} else if v == (curV + 1) {
			c += s.CountPaths(n, counted)
		}
	}

	return c
}

func (s *State) CountAllPaths(rating bool) int {
	tot := 0

	for y, l := range s.Lines {
		for x, r := range l {
			if r == '0' {
				m := map[xy.XY]bool{}
				if rating {
					m = nil
				}

				tot += s.CountPaths(xy.XY{x, y}, m)
			}
		}
	}

	return tot
}

func Solve1(inp string) any {
	s := parseInput(inp)

	return s.CountAllPaths(false)
}

func Solve2(inp string) any {
	s := parseInput(inp)

	return s.CountAllPaths(true)
}
