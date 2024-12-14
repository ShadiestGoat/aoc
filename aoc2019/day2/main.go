package day2

import (
	"maps"

	"github.com/shadiestgoat/aoc/aoc2019/intcode"
)

func Solve1(inp string) any {
	code := intcode.ParseIntCode(inp)
	code[1] = 12
	code[2] = 2

	comp := &intcode.Computer{
		Code:   code,
	}
	comp.RunIntCode()

	return code[0]
}

// brute force, but eh wtv, I don't want to figure out dependency stuff
func Solve2(inp string) any {
	code := intcode.ParseIntCode(inp)

	for i := 0; i <= 99; i++ {
		// I initially tried to do just i-99, but given that these positions can be referenced by other ops, that doesn't work
		for j := 0; j <= 99; j++ {
			tmpCode := maps.Clone(code)

			tmpCode[1] = i
			tmpCode[2] = j

			comp := &intcode.Computer{
				Code:   tmpCode,
			}
			comp.RunIntCode()

			if comp.Code[0] == 19690720 {
				return i * 100 + j 
			}
		}
	}

	return nil
}
