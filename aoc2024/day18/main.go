package day18

import (
	"errors"
	"slices"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/xy"
)

// Are you kidding me? Another path finding puzzle?

func parseInput(inp string) []xy.XY {
	return sparse.SplitAndParseFunc(inp, "\n", func(s string) xy.XY {
		return xy.XYFromArr(sparse.SplitAndParseInt(s, ","))
	})
}

func InitializeGraph(max int, coords []xy.XY) (*dijkstra.MappedGraph[xy.XY], map[xy.XY]bool) {
	badSpots := map[xy.XY]bool{}
	g := dijkstra.NewMappedGraph[xy.XY]()

	for _, c := range coords {
		badSpots[c] = true
	}

	size := xy.XY{max + 1, max + 1}

	for y := 0; y <= max; y++ {
		for x := 0; x <= max; x++ {
			bc := xy.XY{x, y}

			if badSpots[bc] {
				continue
			}

			arcs := map[xy.XY]uint64{}
			for _, d := range xy.ALL_DIRECT_DIRS {
				c := bc.Add(d)
				if c.OutOfBounds(size) || badSpots[c] {
					continue
				}

				arcs[c] = 1
			}

			g.AddEmptyVertex(bc)
			g.AddVertexAndArcs(bc, arcs)
		}
	}

	return &g, badSpots
}

func GenericSolve1(inp string, max int, corruptedAmt int) int {
	coords := parseInput(inp)
	g, _ := InitializeGraph(max, coords[:corruptedAmt])

	best, err := g.Shortest(xy.XY{0, 0}, xy.XY{max, max})
	utils.PanicIfErr(err, "finding best path")

	return int(best.Distance)
}

func GenericSolve2(inp string, max int, baseSafeAmt int) string {
	coords := parseInput(inp)
	g, badSpots := InitializeGraph(max, coords[:baseSafeAmt])
	lastTestXY := xy.XY{}

	offset := 0
	for {
		lastBestPath, err := g.Shortest(xy.XY{}, xy.XY{max, max})
		if err != nil {
			if errors.Is(err, dijkstra.ErrNoPath) {
				break
			}

			utils.PanicIfErr(err, "finding path")
		}

		for {
			nc := coords[baseSafeAmt+offset]
			offset++

			badSpots[nc] = true
			g.RemoveVertexAndArcs(nc)

			if slices.Contains(lastBestPath.Path, nc) {
				lastTestXY = nc
				break
			}
		}
	}

	return lastTestXY.String()
}

func Solve1(inp string) any {
	return GenericSolve1(inp, 70, 1024)
}

func Solve2(inp string) any {
	return GenericSolve2(inp, 70, 1024)
}
