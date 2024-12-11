package day11

import (
	"github.com/shadiestgoat/aoc/utils"
)

func ParseInputArray(inp string) []int {
	return utils.SplitAndParseInt(inp, " ")
}

func RunGameArray(l []int, moves int) int {
	for i := 0; i < moves; i++ {
		cur := []int{}

		for _, v := range l {
			cur = append(cur, numberLogic(v)...)
		}

		l = cur
	}

	return len(l)
}