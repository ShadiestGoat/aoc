package day5

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

const (
	ROW_COUNT = 128
	COL_COUNT = 8
)

func doSearch(min, max int, instructions string, lower rune) int {
	for _, d := range instructions {
		half := (max - min + 1) / 2

		if d == lower {
			max -= half
		} else {
			min += half
		}
	}

	return min
}

// Parses a row input into {row, col}
func ParseRow(r string) [2]int {
	return [2]int{
		doSearch(0, ROW_COUNT-1, r[:7], 'F'),
		doSearch(0, COL_COUNT-1, r[7:], 'L'),
	}
}

func SeatID(v [2]int) int {
	return v[0]*8 + v[1]
}

func parseInput(inp string) utils.Iterator[[2]int] {
	passes := strings.Split(inp, "\n")
	i := 0

	return utils.NewScannerUtil(func() ([2]int, bool) {
		if i >= len(passes) {
			return [2]int{}, false
		}

		v := ParseRow(passes[i])
		i++

		return v, true
	})
}

func Solve1(inp string) any {
	max := 0

	s := parseInput(inp)

	for s.Next() {
		v := s.Value()
		id := SeatID(v)

		if id > max {
			max = id
		}
	}

	return max
}

func Solve2(inp string) any {
	s := parseInput(inp)

	pos := [ROW_COUNT][COL_COUNT]bool{}
	for s.Next() {
		v := s.Value()
		pos[v[0]][v[1]] = true
	}

	for i, row := range pos {
		for j, col := range row {
			if col {
				continue
			}
			if j+1 < COL_COUNT && !row[j+1] {
				continue
			}
			if j > 0 && !row[j-1] {
				continue
			}

			return SeatID([2]int{i, j})
		}
	}

	return nil
}
