package day8

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type State struct {
	Map []string
	// Frequency -> list of coords for of the antennas
	FrequencyCoords map[rune][]utils.XY
}

func (s *State) getAntis(maxPerNodeDirection int, exactMatch bool) int {
	antis := map[utils.XY]bool{}
	bounds := utils.XY{len(s.Map[0]) - 1, len(s.Map) - 1}
	min := 1
	if exactMatch {
		min = 0
	}

	for _, allCoords := range s.FrequencyCoords {
		for i, c1 := range allCoords {
			for _, c2 := range allCoords[i + 1:] {
				diff := c2.Add(c1.Mul(-1))

				for i := min; i <= maxPerNodeDirection; i++ {
					foundAnti := false
					ac1 := c1.Add(diff.Mul(-i))
					ac2 := c2.Add(diff.Mul(i))
	
					if !ac1.OutOfBounds(bounds) {
						foundAnti = true
						antis[ac1] = true
					}
					if !ac2.OutOfBounds(bounds) {
						foundAnti = true
						antis[ac2] = true
					}

					if !foundAnti {
						break
					}
				}
			}
		}
	}

	return len(antis)
}

func parseInput(inp string) *State {
	lines := strings.Split(inp, "\n")
	s := &State{
		Map:             lines,
		FrequencyCoords: map[rune][]utils.XY{},
	}

	for y, l := range lines {
		for x, r := range l {
			if r == '.' {
				continue
			}
			s.FrequencyCoords[r] = append(s.FrequencyCoords[r], utils.XY{x, y})
		}
	}

	return s
}

func Solve1(inp string) any {
	s := parseInput(inp)

	return s.getAntis(1, false)
}

func Solve2(inp string) any {
	s := parseInput(inp)

	return s.getAntis(len(s.Map[0]) * 3, true)
}
