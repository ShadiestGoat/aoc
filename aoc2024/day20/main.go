package day20

import (
	"strings"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/xy"
)

type Track struct {
	Data    []rune
	perLine int
	g       *dijkstra.MappedGraph[xy.XY]

	Start, End xy.XY
}

func (t Track) xyToI(c xy.XY) int {
	return c[1]*t.perLine + c[0]
}

func (t Track) iToXY(i int) xy.XY {
	return xy.XY{
		i % t.perLine,
		i / t.perLine,
	}
}

func (t Track) atCoord(c xy.XY) rune {
	i := t.xyToI(c)
	if i < 0 || i >= len(t.Data) {
		return '#'
	}

	return t.Data[i]
}

func parseInput(inp string) *Track {
	g := dijkstra.NewMappedGraph[xy.XY]()

	iStart, iEnd := strings.Index(inp, "S"), strings.Index(inp, "E")
	perLine := strings.Index(inp, "\n")

	arr := []rune(inp)
	arr[iStart] = '.'
	arr[iEnd] = '.'

	t := &Track{
		Data:    arr,
		perLine: perLine + 1,
		g:       &g,
	}
	t.Start, t.End = t.iToXY(iStart), t.iToXY(iEnd)

	for i, v := range arr {
		if v == '#' || v == '\n' {
			continue
		}

		bc := t.iToXY(i)

		dirs := map[xy.XY]uint64{}

		g.AddEmptyVertex(bc)
		for _, d := range xy.ALL_DIRECT_DIRS {
			nc := bc.Add(d)
			if t.atCoord(nc) == '#' {
				continue
			}

			dirs[nc] = 1
		}

		g.AddVertexAndArcs(bc, dirs)
	}

	return t
}

func GenericSolve(inp string, maxCheatTime int, min int) int {
	t := parseInput(inp)

	shortest, err := t.g.Shortest(t.Start, t.End)
	utils.PanicIfErr(err, "finding path")

	tot := 0
	shortcuts := map[int]int{}

	for i, cur := range shortest.Path {
		for j, target := range shortest.Path[i + 1:] {
			dist := cur.ManhattanDistanceTo(target)
			if dist > maxCheatTime {
				continue
			}

			saved := j + 1 - dist
			if saved < min {
				continue
			}

			tot++
			shortcuts[saved]++
		}
	}

	return tot
}

func Solve1(inp string) any {
	return GenericSolve(inp, 2, 100)
}

func Solve2(inp string) any {
	return GenericSolve(inp, 20, 100)
}
