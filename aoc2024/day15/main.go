package day15

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils/xarr"
	"github.com/shadiestgoat/aoc/utils/xy"
)

type Warehouse struct {
	Pos xy.XY
	Map [][]rune
}

func (w *Warehouse) mapBackup() [][]rune {
	c := make([][]rune, len(w.Map))

	for i, l := range w.Map {
		c[i] = slices.Clone(l)
	}

	return c
}

func (w *Warehouse) tileAt(c xy.XY) rune {
	if c.OutOfBounds(xy.GetSize(w.Map)) {
		return '#'
	}

	return w.Map[c[1]][c[0]]
}

func (w *Warehouse) moveTile(toMove xy.XY, dir xy.XY, noPairRecurse bool) bool {
	tTile := w.tileAt(toMove)
	nCoord := toMove.Add(dir)
	nextTile := w.tileAt(nCoord)

	switch nextTile {
	case '#':
		return false
	case 'O', '[', ']':
		if !w.moveTile(nCoord, dir, false) {
			return false
		}
	}

	if dir[1] != 0 && !noPairRecurse && (tTile == '[' || tTile == ']') {
		x := 1
		if tTile == ']' {
			x = -1
		}

		if !w.moveTile(toMove.Add(xy.XY{x}), dir, true) {
			return false
		}
	}

	w.Map[nCoord[1]][nCoord[0]] = tTile
	w.Map[toMove[1]][toMove[0]] = '.'

	return true
}

func (w *Warehouse) Move(dir xy.XY) {
	next := w.Pos.Add(dir)

	switch w.Map[next[1]][next[0]] {
	case '#':
		return
	case '.':
		w.Pos = next
		return
	}

	backup := w.mapBackup()

	if w.moveTile(next, dir, false) {
		w.Pos = next
	} else {
		w.Map = backup
	}
}

func (w Warehouse) String() string {
	str := ""

	for y, l := range w.Map {
		if w.Pos[1] == y {
			arr := slices.Clone(l)
			arr[w.Pos[0]] = '@'
			l = arr
		}

		str += "\n" + string(l)
	}

	return str[1:]
}

func parseDirs(inp string) []xy.XY {
	return xarr.Map([]rune(strings.Join(strings.Split(inp, "\n"), "")), func(r rune) xy.XY {
		switch r {
		case '^':
			return xy.DIR_UP
		case 'v':
			return xy.DIR_DOWN
		case '<':
			return xy.DIR_LEFT
		case '>':
			return xy.DIR_RIGHT
		}

		return xy.XY{}
	})
}

func parseInput1(inp string) (*Warehouse, []xy.XY) {
	w := &Warehouse{
		Pos: [2]int{},
		Map: [][]rune{},
	}

	spl := strings.Split(inp, "\n\n")

	for y, l := range strings.Split(spl[0], "\n") {
		arr := []rune(l)
		robotIndex := slices.Index(arr, '@')
		if robotIndex != -1 {
			arr[robotIndex] = '.'
			w.Pos = xy.XY{robotIndex, y}
		}

		w.Map = append(w.Map, arr)
	}

	return w, parseDirs(spl[1])
}

func parseInput2(inp string) (*Warehouse, []xy.XY) {
	w := &Warehouse{
		Pos: [2]int{},
		Map: [][]rune{},
	}

	spl := strings.Split(inp, "\n\n")

	for y, l := range strings.Split(spl[0], "\n") {
		arr := []rune(l)
		robotIndex := slices.Index(arr, '@')
		if robotIndex != -1 {
			arr[robotIndex] = '.'
			w.Pos = xy.XY{robotIndex * 2, y}
		}

		curLines := []rune{}
		for _, r := range arr {
			if r == 'O' {
				curLines = append(curLines, '[', ']')
			} else {
				curLines = append(curLines, r, r)
			}

		}

		w.Map = append(w.Map, curLines)
	}

	return w, parseDirs(spl[1])
}

func (w Warehouse) getAnswer() int {
	t := 0

	for y, l := range w.Map {
		for x, v := range l {
			if v == 'O' || v == '[' {
				t += y*100 + x
			}
		}
	}

	return t
}

func Solve1(inp string) any {
	w, dirs := parseInput1(inp)

	for _, d := range dirs {
		w.Move(d)
	}

	return w.getAnswer()
}

func Solve2(inp string) any {
	w, dirs := parseInput2(inp)

	for _, d := range dirs {
		w.Move(d)
	}

	return w.getAnswer()
}
