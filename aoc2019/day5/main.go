package day5

import (
	"github.com/shadiestgoat/aoc/aoc2019/intcode"
)

func getDiagnosticCode(id int, inp string) int {
	comp := &intcode.Computer{
		Input:  []int{id},
		Output: []int{},
		Code: intcode.ParseIntCode(inp),
	}

	comp.RunIntCode()

	dCode := 0
	for _, v := range comp.Output {
		if v == 0 {
			continue
		}

		dCode = v
	}

	return dCode
}

func Solve1(inp string) any {
	return getDiagnosticCode(1, inp)
}

func Solve2(inp string) any {
	return getDiagnosticCode(5, inp)
}
