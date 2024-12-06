package day4

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
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

	ALL_PART_2_CHECKS = [4]map[utils.XY]rune{}
)

func init() {
	ALL_PART_2_CHECKS[0] = map[utils.XY]rune{
		{-1, -1}: 'M',
		{1, -1}:  'M',
		{-1, 1}:  'S',
		{1, 1}:   'S',
	}
	
	// Mess? What mess?
	for i := 1; i < 4; i++ {
		curCheck := map[utils.XY]rune{}

		for dir, c := range ALL_PART_2_CHECKS[i - 1] {
			var newDir utils.XY

			if dir[0] == dir[1] {
				newDir = utils.XY{dir[0] * -1, dir[1]}
			} else {
				newDir = utils.XY{dir[0], dir[0]}
			}

			curCheck[newDir] = c
		}

		ALL_PART_2_CHECKS[i] = curCheck
	}
}

type Board []string

func parseInput(inp string) Board {
	return strings.Split(inp, "\n")
}

func (b Board) checkCoord(xy utils.XY, t rune) bool {
	if xy[0] < 0 || xy[1] < 0 || xy[1] >= len(b) || xy[0] >= len(b[xy[1]]) {
		return false
	}

	return rune(b[xy[1]][xy[0]]) == t
}

func (b Board) searchDir(xy utils.XY, dir utils.XY, targetWord string) bool {
	// 0 is already confirmed
	targetI := 1

	for {
		if !b.checkCoord(xy.Add(dir.Mul(targetI)), rune(targetWord[targetI])) {
			return false
		}

		targetI++

		if targetI == len(targetWord) {
			return true
		}
	}
}

func (b Board) dirChecks(xy utils.XY, checks map[utils.XY]rune) bool {
	for dir, t := range checks {
		if !b.checkCoord(xy.Add(dir), t) {
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
				if b.dirChecks(utils.XY{j, i}, check) {
					tot++
					break
				}
			}
		}
	}

	return tot
}
