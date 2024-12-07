package day7

import (
	"math"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Op int

const (
	OP_MUL Op = iota
	OP_ADD
	OP_CONCAT
)

func (o Op) Do(a, b int) int {
	switch o {
	case OP_MUL:
		return a * b
	case OP_ADD:
		return a + b
	case OP_CONCAT:
		size := int(math.Floor(math.Log10(float64(b)))) + 1

		return int(math.Pow10(size)*float64(a)) + b
	}

	return 0
}

type Equation struct {
	Res  int
	Vals []int
}

func (e *Equation) possibilityRecursor(cur int, valIndex int, ops []Op) bool {
	if valIndex >= len(e.Vals) {
		return e.Res == cur
	}

	rhs := e.Vals[valIndex]
	for _, op := range ops {
		v := op.Do(cur, rhs)
		if v > e.Res {
			continue
		}

		if e.possibilityRecursor(v, valIndex+1, ops) {
			return true
		}
	}

	return false
}

func (e *Equation) CanBePossible(ops []Op) bool {
	return e.possibilityRecursor(e.Vals[0], 1, ops)
}

func parseInput(inp string) []*Equation {
	return utils.SplitAndParseFunc(inp, "\n", func(s string) *Equation {
		spl := strings.Split(s, ": ")

		return &Equation{
			Res:  utils.ParseInt(spl[0]),
			Vals: utils.SplitAndParseInt(spl[1], " "),
		}
	})
}

func GenericSolve(inp string, ops []Op) int {
	eqs := parseInput(inp)
	tot := 0

	for _, e := range eqs {
		if e.CanBePossible(ops) {
			tot += e.Res
		}
	}

	return tot
}

func Solve1(inp string) any {
	return GenericSolve(inp, []Op{OP_MUL, OP_ADD})
}

func Solve2(inp string) any {
	return GenericSolve(inp, []Op{OP_MUL, OP_ADD, OP_CONCAT})
}
