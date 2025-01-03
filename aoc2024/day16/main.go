package day16

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/xy"

	// In case you're wandering, no I'm not happy about using this lib. But also, I cannot be fucked to implement this rn
	"github.com/RyanCarrier/dijkstra/v2"
)

func makeGraph(m []string, startPos, startDir, endPos xy.XY) *xRoadState {
	// We will use a 4 vertex system for connections - 1 for each allowed dir, but none that are connected directly
	/*
	       C
	       |
	       ^
	   A--< >---B
	       v
	       |
	   	D

	   (+ A<->B, C<->D)

	   Where they can all go to each other, but not via some intermediate point. So A -> M -> D is impossible, only A->D.
	*/
	size := xy.GetSizeString(m)
	crossRoads := []xy.XY{startPos, endPos}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			bc := xy.XY{x, y}
			if m[bc[1]][bc[0]] == '#' {
				continue
			}

			// coordIndex[bc] = map[xy.XY]int{}
			founds := [4]bool{}

			for i, d := range xy.ALL_DIRECT_DIRS {
				c := bc.Add(d)
				if c.OutOfBounds(size) || m[c[1]][c[0]] == '#' {
					continue
				}
				founds[i] = true
			}

			isCrosser := false
			for i, v := range append(founds[1:], founds[0]) {
				if founds[i] && v {
					isCrosser = true
					break
				}
			}

			if isCrosser && (bc != startPos && bc != endPos) {
				crossRoads = append(crossRoads, bc)
			}
		}
	}

	s := &xRoadState{
		g:        &dijkstra.Graph{},
		ref:      map[xy.XY]map[xy.XY]int{},
		revRef:   map[int][2]xy.XY{},
		allRoads: crossRoads,
		m:        m,
		size:     size,
	}

	for _, c := range crossRoads {
		for _, d := range xy.ALL_DIRECT_DIRS {
			s.parse(c, d)
		}
	}

	s.getIndex(startPos, startDir)

	for _, dirs := range s.ref {
		for _, d := range xy.ALL_DIRECT_DIRS {
			v, ok := dirs[d]
			if !ok {
				continue
			}

			if v2, ok2 := dirs[d.RotateVector(4)]; ok2 {
				s.g.AddArc(v, v2, 1)
			}

			if v2, ok2 := dirs[d.RotateVector(2)]; ok2 {
				s.g.AddArc(v, v2, 1001)
			}

			if v2, ok2 := dirs[d.RotateVector(-2)]; ok2 {
				s.g.AddArc(v, v2, 1001)
			}
		}
	}

	return s
}

type xRoadState struct {
	g        *dijkstra.Graph
	ref      map[xy.XY]map[xy.XY]int
	revRef   map[int][2]xy.XY
	allRoads []xy.XY
	m        []string
	size     xy.XY
}

func (s *xRoadState) getIndex(c xy.XY, dir xy.XY) int {
	if s.ref[c] == nil {
		i := s.g.AddNewEmptyVertex()
		s.ref[c] = make(map[xy.XY]int)
		s.ref[c][dir] = i
		s.revRef[i] = [2]xy.XY{c, dir}
	} else {
		if _, ok := s.ref[c][dir]; !ok {
			i := s.g.AddNewEmptyVertex()
			s.ref[c][dir] = i
			s.revRef[i] = [2]xy.XY{c, dir}
		}
	}

	return s.ref[c][dir]
}

func (s *xRoadState) parse(c xy.XY, dir xy.XY) {
	i := 1

	for {
		nc := c.Add(dir.Mul(i))
		// If we hit a wall before we hit a crossroad, it means we got ticked: its a dead end
		if nc.OutOfBounds(s.size) || s.m[nc[1]][nc[0]] == '#' {
			return
		}
		if !slices.Contains(s.allRoads, nc) {
			i++
			continue
		}

		gi1, gi2 := s.getIndex(c, dir), s.getIndex(nc, dir.RotateVector(-4))
		s.g.AddArc(gi1, gi2, uint64(i)-1)

		return
	}
}

func parseInput(inp string) (dir xy.XY, m []string, posS, posE xy.XY, gd *xRoadState) {
	sI := strings.Index(inp, "S")
	eI := strings.Index(inp, "E")
	arr := []rune(inp)
	arr[sI] = '.'
	arr[eI] = '.'

	lines := strings.Split(string(arr), "\n")

	perLine := len(lines) + 1
	posS = xy.XY{sI % perLine, sI / perLine}
	posE = xy.XY{eI % perLine, eI / perLine}

	return xy.DIR_RIGHT, lines, posS, posE, makeGraph(lines, posS, xy.DIR_RIGHT, posE)
}

func Solve1(inp string) any {
	sDir, _, start, end, gd := parseInput(inp)

	startIndex := gd.ref[start][sDir]
	best := dijkstra.BestPath[int]{}

	for _, finIndex := range gd.ref[end] {
		v, err := gd.g.Shortest(startIndex, finIndex)
		utils.PanicIfErr(err, "finding shortest")

		if best.Distance == 0 || v.Distance < best.Distance {
			best = v
		}
	}

	return int(best.Distance)
}

func path(c1, c2 xy.XY) []xy.XY {
	dir := c2.Add(c1.Mul(-1)).Unit()
	o := []xy.XY{c1}
	last := c1

	for {
		n1 := last.Add(dir)
		o = append(o, n1)
		if n1 == c2 {
			return o
		}

		last = n1
	}
}

func Solve2(inp string) any {
	sDir, _, start, end, gd := parseInput(inp)

	startIndex := gd.ref[start][sDir]
	best := uint64(0)
	bestPaths := map[xy.XY]bool{}

	for _, finIndex := range gd.ref[end] {
		v, err := gd.g.ShortestAll(startIndex, finIndex)
		utils.PanicIfErr(err, "finding shortest")

		if best == 0 || v.Distance <= best {
			if v.Distance != best {
				bestPaths = make(map[xy.XY]bool)
				best = v.Distance
			}

			for _, p := range v.Paths {
				for i, cur := range p[1:] {
					lastCoord := gd.revRef[p[i]][0]
					curCoord := gd.revRef[cur][0]

					for _, coord := range path(lastCoord, curCoord) {
						bestPaths[coord] = true
					}
				}
			}
		}
	}

	return len(bestPaths)
}
