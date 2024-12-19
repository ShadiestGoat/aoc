package day18

import (
	"errors"
	"slices"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/shadiestgoat/aoc/utils"
)

// Are you kidding me? Another path finding puzzle?

func parseInput(inp string) []utils.XY {
	return utils.SplitAndParseFunc(inp, "\n", func(s string) utils.XY {
		return utils.XYFromArr(utils.SplitAndParseInt(s, ","))
	})
}

func InitializeGraph(max int, coords []utils.XY) (*dijkstra.MappedGraph[utils.XY], map[utils.XY]bool) {
	badSpots := map[utils.XY]bool{}
	g := dijkstra.NewMappedGraph[utils.XY]()

	for _, c := range coords {
		badSpots[c] = true
	}

	size := utils.XY{max + 1, max + 1}

	for y := 0; y <= max; y++ {
		for x := 0; x <= max; x++ {
			bc := utils.XY{x, y}

			if badSpots[bc] {
				continue
			}

			arcs := map[utils.XY]uint64{}
			for _, d := range utils.ALL_DIRECT_DIRS {
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
	
	best, err := g.Shortest(utils.XY{0, 0}, utils.XY{max, max})
	utils.PanicIfErr(err, "finding best path")

	return int(best.Distance)
}

func GenericSolve2(inp string, max int, baseSafeAmt int) string {
	coords := parseInput(inp)
	g, badSpots := InitializeGraph(max, coords[:baseSafeAmt])
	lastTestXY := utils.XY{}

	offset := 0
	for {
		lastBestPath, err := g.Shortest(utils.XY{}, utils.XY{max, max})
		if err != nil {
			if errors.Is(err, dijkstra.ErrNoPath) {
				break
			}
			
			utils.PanicIfErr(err, "finding path")
		}

		for {
			nc := coords[baseSafeAmt + offset]
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
