package day15

import (
	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/shadiestgoat/aoc/aoc2019/intcode"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/xarr"
	"github.com/shadiestgoat/aoc/utils/xprint"
	"github.com/shadiestgoat/aoc/utils/xy"
)

var dirToI = map[xy.XY]int{
	xy.DIR_UP:    1,
	xy.DIR_DOWN:  2,
	xy.DIR_LEFT:  3,
	xy.DIR_RIGHT: 4,
}

func draw(pos xy.XY, walls map[xy.XY]bool) {
	min, max := xy.MinMaxOfCoords(append(xarr.MapKeys(walls), pos))

	width := max[0] - min[0]

	str := ""
	for y := min[1]; y <= max[1]; y++ {
		r := make([]rune, width+1)

		for i, x := 0, min[0]; x <= max[0]; x, i = x+1, i+1 {
			if walls[xy.XY{x, y}] {
				r[i] = '#'
			} else {
				r[i] = ' '
			}
		}

		if y == pos[1] {
			r[pos[0]-min[0]] = 'O'
		}

		str += "\n" + string(r)
	}

	xprint.ClearAndPrint(str[1:])
}

type MoveFunc = func(dir xy.XY) int

func findMap(f MoveFunc) (map[xy.XY]bool, xy.XY) {
	walls := map[xy.XY]bool{}

	curPos := xy.XY{}
	dir := xy.DIR_UP
	endPos := xy.XY{}

	next := false
	for {
		np := curPos.Add(dir)

		o := f(dir)

		if o == 0 {
			walls[np] = true
			dir = dir.RotateUnitVector(2)

			continue
		} else if o == 2 {
			endPos = np
		}

		curPos = np
		if next {
			return walls, endPos
		}

		dir = dir.RotateUnitVector(-2)
		if np.IsAtOrigin() && !endPos.IsAtOrigin() {
			next = true
		}
	}
}

func GenericSolve1(f MoveFunc) int {
	walls, endPos := findMap(f)

	g := dijkstra.NewMappedGraph[xy.XY]()

	min, max := xy.MinMaxOfCoords(append(xarr.MapKeys(walls), endPos, xy.XY{}))
	min = min.Add(xy.XY{-1, -1})
	max = max.Add(xy.XY{1, 1})

	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			c := xy.XY{x, y}
			if walls[c] {
				continue
			}

			arcs := map[xy.XY]uint64{}
			for _, d := range xy.ALL_DIRECT_DIRS {
				if walls[c.Add(d)] {
					continue
				}
				arcs[c.Add(d)] = 1
			}

			g.AddEmptyVertex(c)
			g.AddVertexAndArcs(c, arcs)
		}
	}

	p, err := g.Shortest(xy.XY{}, endPos)
	utils.PanicIfErr(err, "searching for path")

	return int(p.Distance)
}

func GenericSolve2(f MoveFunc) int {
	walls, endPos := findMap(f)

	curs := []xy.XY{endPos}

	i := 0
	for len(curs) != 0 {
		newCurs := []xy.XY{}

		for _, c := range curs {
			for _, d := range xy.ALL_DIRECT_DIRS {
				nc := c.Add(d)
				if walls[nc] {
					continue
				}

				walls[nc] = true
				newCurs = append(newCurs, nc)
			}
		}

		curs = newCurs
		if len(newCurs) != 0 {
			i++
		}
	}

	return i
}

func genericSolveX(inp string, solver func(MoveFunc) int) int {
	comp := intcode.NewComp(inp)

	return solver(func(dir xy.XY) int {
		comp.Input = []int{dirToI[dir]}
		comp.RunIntCode()
		return comp.ConsumeOutput()[0]
	})
}

func Solve1(inp string) any {
	return genericSolveX(inp, GenericSolve1)
}

func Solve2(inp string) any {
	return genericSolveX(inp, GenericSolve2)
}
