package day12

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type State struct {
	Regions map[int][]utils.XY
	ParsedMap [][]int
}

func (s *State) regionAt(xy utils.XY) int {
	if xy.OutOfBounds(utils.GetSize(s.ParsedMap)) {
		return -1
	}

	return s.ParsedMap[xy[1]][xy[0]]
}

func (s *State) GetEdgeAmount(r int) int {
	c := 0

	for _, coord := range s.Regions[r] {
		for _, dir := range utils.ALL_DIRECT_DIRS {
			if s.regionAt(coord.Add(dir)) != r {
				c++
			}
		}
	}

	return c
}

func (s *State) GetStartingEdgeAmount(r int) int {
	c := 0

	for _, coord := range s.Regions[r] {
		for _, dir := range utils.ALL_DIRECT_DIRS {
			if s.regionAt(coord.Add(dir)) == r {
				continue
			}

			if s.regionAt(dir.RotateUnitVector(2).Add(coord)) == r {
				if s.regionAt(dir.RotateUnitVector(1).Add(coord)) != r {
					continue
				}
			}

			c++
		}
	}

	return c
}

func parseInput(inp string) (*State, map[int]rune) {
	regionInfo := map[int]rune{}
	s := &State{
		Regions:   map[int][]utils.XY{},
		ParsedMap: [][]int{},
	}

	lines := strings.Split(inp, "\n")

	checkRegion := func (x, y int, r rune) int {
		lrID := s.ParsedMap[y][x]
		if regionInfo[lrID] == r {
			return lrID
		}

		return -1
	}

	maxRegion := -1
	for y, l := range lines {
		s.ParsedMap = append(s.ParsedMap, []int{})

		for x, r := range l {
			lastRegion := -1
			if x != 0 {
				lastRegion = checkRegion(x - 1, y, r)
			}

			if lastRegion == -1 && y > 0 {
				tmp := utils.XY{x, y}
				for !tmp.OutOfBounds(utils.GetSizeString(lines)) {
					if r == rune(l[tmp[0]]) {
						if r == rune(lines[y - 1][tmp[0]]) {
							lastRegion = checkRegion(tmp[0], y - 1, r)
							break
						}
					} else {
						break
					}

					tmp = tmp.Add(utils.XY{1, 0})
				}
			}

			if lastRegion == -1 {
				maxRegion++
				lastRegion = maxRegion
				regionInfo[maxRegion] = r
			}

			s.ParsedMap[y] = append(s.ParsedMap[y], lastRegion)
			s.Regions[lastRegion] = append(s.Regions[lastRegion], utils.XY{x, y})
		}
	}

	// So, in theory, there can be regions that are above, but not connected to the rest.
	// See subtest 4 of Solve1
	foundSmt := true
	for foundSmt {
		foundSmt = false

		for y, l := range lines {
			if y == len(lines) - 1 {
				break
			}

			for x, r := range l {
				if r != rune(lines[y + 1][x]) {
					continue
				}

				cRID := s.ParsedMap[y][x]
				// The new region id
				rID := s.ParsedMap[y + 1][x]
				if rID == cRID {
					continue
				}

				// Old regions, time to merge :3
				for _, c := range s.Regions[cRID] {
					s.ParsedMap[c[1]][c[0]] = rID
					s.Regions[rID] = append(s.Regions[rID], c)
				}

				delete(s.Regions, cRID)

				foundSmt = true
				break
			}

			if foundSmt {
				break
			}
		}
	}

	return s, regionInfo
}

func Solve1(inp string) any {
	s, _ := parseInput(inp)

	price := 0
	for r, coords := range s.Regions {
		edges := s.GetEdgeAmount(r)

		price += edges * len(coords)
	}

	return price
}

func Solve2(inp string) any {
	s, _ := parseInput(inp)

	price := 0
	for r, coords := range s.Regions {
		edges := s.GetStartingEdgeAmount(r)

		price += edges * len(coords)
	}

	return price
}
