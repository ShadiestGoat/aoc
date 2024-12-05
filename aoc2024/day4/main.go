package day4

import (
	"strings"
)

const (
	TARGET_WORD_1 = "XMAS"
	TARGET_WORD_2 = "MAS"
)

var (
	ALL_DIRS = [][2]int{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	ALL_PART_2_CHECKS = [4]map[XY]rune{}
)

func init() {
	ALL_PART_2_CHECKS[0] = map[XY]rune{
		{-1, -1}: 'M',
		{1, -1}:  'M',
		{-1, 1}:  'S',
		{1, 1}:   'S',
	}
	
	// Mess? What mess?
	for i := 1; i < 4; i++ {
		curCheck := map[XY]rune{}

		for dir, c := range ALL_PART_2_CHECKS[i - 1] {
			var newDir XY

			if dir[0] == dir[1] {
				newDir = XY{dir[0] * -1, dir[1]}
			} else {
				newDir = XY{dir[0], dir[0]}
			}

			curCheck[newDir] = c
		}

		ALL_PART_2_CHECKS[i] = curCheck
	}
}

type Board []string
type XY [2]int

func (c XY) add(c2 XY) XY {
	return [2]int{c[0] + c2[0], c[1] + c2[1]}
}

func (c XY) mulCoord(c2 XY) XY {
	return [2]int{c[0] * c2[0], c[1] * c2[1]}
}

func (c XY) mul(v int) XY {
	return [2]int{c[0] * v, c[1] * v}
}

func parseInput(inp string) Board {
	return strings.Split(inp, "\n")
}

func (b Board) checkCoord(xy XY, t rune) bool {
	if xy[0] < 0 || xy[1] < 0 || xy[1] >= len(b) || xy[0] >= len(b[xy[1]]) {
		return false
	}

	return rune(b[xy[1]][xy[0]]) == t
}

func (b Board) searchDir(xy XY, dir XY, targetWord string) bool {
	// 0 is already confirmed
	targetI := 1

	for {
		if !b.checkCoord(xy.add(dir.mul(targetI)), rune(targetWord[targetI])) {
			return false
		}

		targetI++

		if targetI == len(targetWord) {
			return true
		}
	}
}

func (b Board) dirChecks(xy XY, checks map[XY]rune) bool {
	for dir, t := range checks {
		if !b.checkCoord(xy.add(dir), t) {
			return false
		}
	}

	return true
}

func Solve1(inp string) any {
	b := parseInput(inp)
	tot := 0

	for i, r := range b {
		for j, l := range r {
			if l != rune(TARGET_WORD_1[0]) {
				continue
			}

			for _, dir := range ALL_DIRS {
				if b.searchDir([2]int{j, i}, dir, TARGET_WORD_1) {
					tot++
				}
			}
		}
	}

	return tot
}

// Its cheating a bit but sshh
// Were basically going to check if char == a, then check around it to see if everything matches correctly
func Solve2(inp string) any {
	b := parseInput(inp)
	tot := 0

	for i, r := range b {
		for j, l := range r {
			if l != 'A' {
				continue
			}

			for _, check := range ALL_PART_2_CHECKS {
				if b.dirChecks(XY{j, i}, check) {
					tot++
					break
				}
			}
		}
	}

	return tot
}
