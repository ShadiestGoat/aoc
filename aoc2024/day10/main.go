package day10

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

var (
	ALL_DIRS = []utils.XY{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

type State struct {
	Lines []string
	size  utils.XY
}

func parseInput(inp string) *State {
	lines := strings.Split(inp, "\n")
	return &State{
		Lines: lines,
		size:  utils.XY{len(lines[0]) - 1, len(lines) - 1},
	}
}

func (s *State) valAt(c utils.XY) rune {
	return rune(s.Lines[c[1]][c[0]])
}

func (s *State) CountPaths(cur utils.XY, counted map[utils.XY]bool) int {
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
	m := map[utils.XY]bool{}
	if rating {
		m = nil
	}

	for y, l := range s.Lines {
		for x, r := range l {
			if r == '0' {
				tot += s.CountPaths(utils.XY{x, y}, m)
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
