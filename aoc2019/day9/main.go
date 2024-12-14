package day9

import "github.com/shadiestgoat/aoc/aoc2019/intcode"

func Solve1(inp string) any {
	return intcode.QuickRun(intcode.ParseIntCode(inp), []int{1})
}

func Solve2(inp string) any {
	return intcode.QuickRun(intcode.ParseIntCode(inp), []int{2})
}
