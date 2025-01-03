package day6

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/sillyset/set3"
	"github.com/shadiestgoat/aoc/utils/xy"
)

type State struct {
	CurPos xy.XY
	OgPos  xy.XY
	Dir    xy.XY

	size       xy.XY
	WallsByCol []set3.Set
	WallsByRow []set3.Set
}

func parseInput(inp string) *State {
	rawPos := strings.Index(inp, "^")
	// Its been a tiring week... I'm not gonna bother with a fully string board like aoc2020/day11
	lines := strings.Split(strings.Replace(inp, "^", ".", 1), "\n")
	perLineChars := len(lines[0]) + 1
	pos := xy.XY{rawPos % perLineChars, rawPos / perLineChars}

	obsByCol := make([]set3.Set, len(lines[0]))
	obsByRow := make([]set3.Set, len(lines))

	for y, l := range lines {
		for x, v := range l {
			if v == '#' {
				obsByCol[x] = obsByCol[x].UpdateAt(uint64(y), 1)
				obsByRow[y] = obsByRow[y].UpdateAt(uint64(x), 1)
			}
		}
	}

	return &State{
		CurPos:     pos,
		OgPos:      pos,
		Dir:        xy.DIR_UP,
		WallsByCol: obsByCol,
		WallsByRow: obsByRow,
		size:       xy.GetSizeString(lines),
	}
}

func (s *State) GoUntilExit() map[xy.XY]bool {
	pos := map[xy.XY]bool{}

	for {
		nextV := s.CurPos.Add(s.Dir)
		if nextV.OutOfBounds(s.size) {
			return pos
		}

		if s.WallsByCol[nextV[0]].At(uint64(nextV[1])) {
			s.Dir = s.Dir.RotateVector(2)

			continue
		}

		s.CurPos = nextV
		pos[s.CurPos] = true
	}
}

func Solve1(inp string) any {
	return len(parseInput(inp).GoUntilExit())
}

func (s *State) AddWall(c xy.XY, v uint64) {
	s.WallsByCol[c[0]] = s.WallsByCol[c[0]].UpdateAt(uint64(c[1]), v)
	s.WallsByRow[c[1]] = s.WallsByRow[c[1]].UpdateAt(uint64(c[0]), v)
}

func (s State) DoesLoop() bool {
	curPos := s.OgPos
	curDir := xy.DIR_UP

	hist := map[[2]xy.XY]bool{
		{curPos, curDir}: true,
	}

	for {
		var walls set3.Set
		var i int

		if curDir[0] == 0 {
			// Going vertically
			walls = s.WallsByCol[curPos[0]]
			i = 1
		} else {
			// Going horizontally
			walls = s.WallsByRow[curPos[1]]
			i = 0
		}

		var f int
		if curDir[i] < 0 {
			f = walls.ClosestSmall(curPos[i])
		} else {
			f = walls.ClosestBig(curPos[i])
		}
		if f == -1 {
			return false
		}

		curPos[i] = f - curDir[i]
		curDir = curDir.RotateVector(2)
		k := [2]xy.XY{curPos, curDir}

		if hist[k] {
			return true
		}
		hist[k] = true
	}
}

func Solve2(inp string) any {
	s := parseInput(inp)
	hist := s.GoUntilExit()

	amt := 0
	for c := range hist {
		s.AddWall(c, 1)

		if s.DoesLoop() {
			amt++
		}

		s.AddWall(c, 0)
	}

	return amt
}
