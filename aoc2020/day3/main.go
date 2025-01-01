package day3

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/funiter"
)

const (
	EMPTY = '.'
	TREE  = '#'
)

type SledPath struct {
	Map string

	slopeX    int
	slopeY    int
	lineSize  int
	lineCount int
	i         int
}

func (s *SledPath) Next() bool {
	s.i++

	return s.i*s.slopeY < s.lineCount
}

func (s *SledPath) Value() byte {
	baseI := (s.i * s.slopeX) % s.lineSize

	return s.Map[baseI+s.i*(s.lineSize+1)*s.slopeY]
}

func parseInput(inp string, slopeX, slopeY int) funiter.Iterator[byte] {
	inp = strings.TrimSpace(inp)

	return &SledPath{
		Map:       inp,
		slopeX:    slopeX,
		slopeY:    slopeY,
		lineSize:  strings.Index(inp, "\n"),
		lineCount: strings.Count(inp, "\n") + 1,
	}
}

func Solve1(inp string) any {
	s := parseInput(inp, 3, 1)
	treeCount := 0

	for s.Next() {
		if s.Value() == TREE {
			treeCount++
		}
	}

	return treeCount
}

func Solve2(inp string) any {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	resp := 1

	for _, cfg := range slopes {
		s := parseInput(inp, cfg[0], cfg[1])
		amt := 0

		for s.Next() {
			if s.Value() == TREE {
				amt++
			}
		}

		resp *= amt
	}

	return resp
}
