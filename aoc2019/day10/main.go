package day10

import (
	"math"
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils/mutils"
	"github.com/shadiestgoat/aoc/utils/xarr"
	"github.com/shadiestgoat/aoc/utils/xy"
)

func parseInput(inp string) *World {
	d := xarr.Map(strings.Split(inp, "\n"), func(s string) []rune {
		return []rune(s)
	})

	return &World{
		Data: d,
		Size: xy.GetSize(d),
	}
}

type World struct {
	Data [][]rune
	Size xy.XY
}

func (w World) countFrom(og xy.XY) (int, []xy.XY) {
	seen := 0
	testedSlopes := map[xy.XY]bool{}

	testOffset := func(x, y int) {
		d := mutils.GCD(x, y)

		off := xy.XY{x / d, y / d}
		if testedSlopes[off] {
			return
		}
		testedSlopes[off] = true

		i := 1
		for {
			c := og.Add(off.Mul(i))
			if c.OutOfBounds(w.Size) {
				return
			}

			if w.Data[c[1]][c[0]] == '#' {
				seen++
				return
			}

			i++
		}
	}

	for y := -len(w.Data) - 1; y <= len(w.Data); y++ {
		for x := -len(w.Data[0]) - 1; x <= len(w.Data[0]); x++ {
			if x == 0 && y == 0 {
				continue
			}

			testOffset(x, y)
		}
	}

	return seen, xarr.MapKeys(testedSlopes)
}

func (w World) bestPlace() (int, xy.XY, []xy.XY) {
	max := 0
	bestCoord := xy.XY{}
	bestSlopes := []xy.XY{}

	for y, row := range w.Data {
		for x, r := range row {
			if r == '#' {
				coord := xy.XY{x, y}
				if c, slopes := w.countFrom(coord); c > max {
					max = c
					bestCoord = coord
					bestSlopes = slopes
				}
			}
		}
	}

	return max, bestCoord, bestSlopes
}

func (w World) findClosestInDir(c, dir xy.XY) (xy.XY, bool) {
	m := 1

	for {
		nc := c.Add(dir.Mul(m))
		if nc.OutOfBounds(w.Size) {
			return xy.XY{}, false
		}

		if w.Data[nc[1]][nc[0]] == '#' {
			return nc, true
		}

		m++
	}
}

func Solve1(inp string) any {
	w := parseInput(inp)

	amt, _, _ := w.bestPlace()

	return amt
}

func Solve2(inp string) any {
	w := parseInput(inp)

	_, coord, offsets := w.bestPlace()
	slices.SortStableFunc(offsets, func(a, b xy.XY) int {
		qa, qb := a.Quadrant(), b.Quadrant()
		if qa < qb {
			return -1
		} else if qa > qb {
			return 1
		}

		aa, ab := math.Tanh(float64(a[0])/-float64(a[1])), math.Tanh(float64(b[0])/-float64(b[1]))
		if aa == ab {
			return 0
		} else if aa < ab {
			return -1
		}

		return 1
	})

	offI := 0
	destroyed := 0
	for {
		nc, ok := w.findClosestInDir(coord, offsets[offI])
		if ok {
			destroyed++
			w.Data[nc[1]][nc[0]] = '.'
			if destroyed == 199 {
				return nc[0]*100 + nc[1]
			}
		}

		offI++
		if offI >= len(offsets) {
			offI = 0
		}
	}
}
