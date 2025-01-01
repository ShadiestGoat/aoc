package day15_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day15"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT_1 = `
#####
##..#
#..##
#####
`

const INPUT_2 = `
 ##   
#..## 
#.#..#
#...# 
 ###  
`

func mover(inp string, s, e utils.XY) day15.MoveFunc {
	inp = strings.Trim(inp, "\n")

	perLine := strings.Index(inp, "\n") + 1

	curCoord := s
	endCoord := e

	return func(dir utils.XY) int {
		nc := curCoord.Add(dir)
		i := nc[1] * perLine + nc[0]

		if i < 0 || i > len(inp) {
			return 0
		}
		if inp[i] != '.' {
			return 0
		}

		curCoord = nc
		if nc == endCoord {
			return 2
		}

		return 1
	}
}

func TestGenericSolve1(t *testing.T) {
	m := mover(INPUT_1, utils.XY{2, 1}, utils.XY{1, 2})
	r := day15.GenericSolve1(m)

	tutils.Assert(t, 2, r)
}

func TestGenericSolve2(t *testing.T) {
	m := mover(INPUT_2, utils.XY{1, 1}, utils.XY{2, 3})
	r := day15.GenericSolve2(m)

	tutils.Assert(t, 4, r)
}
