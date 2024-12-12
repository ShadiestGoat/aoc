package day2

import "github.com/shadiestgoat/aoc/utils"

func ParseInput(inp string) []int {
	return utils.SplitAndParseInt(inp, ",")
}

func doOp(inp []int, pos int) {
	var (
		op     = inp[pos]
		v1     = inp[inp[pos+1]]
		v2     = inp[inp[pos+2]]
		resPos = inp[pos+3]
	)

	switch op {
	case 1:
		inp[resPos] = v1 + v2
	case 2:
		inp[resPos] = v1 * v2
	default:
		panic("Bad OpCode")
	}
}

func RunIntCode(code []int) {
	pos := 0
	for code[pos] != 99 {
		doOp(code, pos)
		pos += 4
	}
}

func Solve1(inp string) any {
	code := ParseInput(inp)
	code[1] = 12
	code[2] = 2

	RunIntCode(code)

	return code[0]
}

// brute force, but eh wtv, I don't want to figure out dependency stuff
func Solve2(inp string) any {
	for i := 0; i <= 99; i++ {
		// I initially tried to do just i-99, but given that these positions can be referenced by other ops, that doesn't work
		for j := 0; j <= 99; j++ {
			code := ParseInput(inp)
			code[1] = i
			code[2] = j

			RunIntCode(code)
			if code[0] == 19690720 {
				return i * 100 + j 
			}
		}
	}

	return nil
}
