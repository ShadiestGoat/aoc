package day18

import (
	"strconv"
	"strings"

	"github.com/shadiestgoat/aoc/utils/sparse"
)

func evalBraces(line *string, doMath func(string) int) bool {
	i1 := strings.LastIndex(*line, "(")
	if i1 == -1 {
		return false
	}

	i2 := strings.Index((*line)[i1:], ")") + i1

	res := doMath((*line)[i1+1 : i2])

	replaceStr(line, strconv.Itoa(res), i1, i2)

	return true
}

func replaceStr(str *string, val string, i1, i2 int) {
	*str = (*str)[:i1] + val + (*str)[i2+1:]
}

func DoMathLTR(line string) int {
	for evalBraces(&line, DoMathLTR) {
	}

	line = strings.ReplaceAll(line, " ", "")

	i := strings.IndexAny(line, "+*")
	cur := sparse.ParseInt(line[:i])
	line = line[i:]

	for {
		if len(line) == 0 {
			return cur
		}

		nextIndex := strings.IndexAny(line[1:], "+*")
		if nextIndex == -1 {
			nextIndex = len(line)
		} else {
			nextIndex++
		}

		nextNum := sparse.ParseInt(line[1:nextIndex])

		if line[0] == '*' {
			cur *= nextNum
		} else {
			cur += nextNum
		}

		line = line[nextIndex:]
	}
}

func mathPlusParser(line *string, sym string, op func(a, b int) int) {
	for strings.Count(*line, sym) != 0 {
		i := strings.Index(*line, sym)

		iMin := strings.LastIndexAny((*line)[:i], "+*")

		iMax := strings.IndexAny((*line)[i+1:], "+*")
		if iMax == -1 {
			iMax = len(*line)
		} else {
			iMax += i + 1
		}

		replaceStr(line, strconv.Itoa(
			op(
				sparse.ParseInt((*line)[iMin+1:i]),
				sparse.ParseInt((*line)[i+1:iMax]),
			),
		), iMin+1, iMax-1)
	}
}

func DoMathPlus(line string) int {
	for evalBraces(&line, DoMathPlus) {
	}
	line = strings.ReplaceAll(line, " ", "")

	mathPlusParser(&line, "+", func(a, b int) int {
		return a + b
	})

	mathPlusParser(&line, "*", func(a, b int) int {
		return a * b
	})

	return sparse.ParseInt(line)
}

func solve(inp string, doMath func(string) int) int {
	sum := 0

	for _, l := range strings.Split(inp, "\n") {
		sum += doMath(l)
	}

	return sum
}

func Solve1(inp string) any {
	return solve(inp, DoMathLTR)
}

func Solve2(inp string) any {
	return solve(inp, DoMathPlus)
}
