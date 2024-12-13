package intcode

import (
	"math"
	"strconv"

	"github.com/shadiestgoat/aoc/utils"
)

var (
	// op -> param count
	param_amounts = map[int]int{
		1: 3,
		2: 3,
		3: 1,
		4: 1,
		5: 2,
		6: 2,
		7: 3,
		8: 3,
	}
	// op -> true, if op writes to last param's addr (tldr: override mechanic)
	write_ops = map[int]bool{
		1: true,
		2: true,
		3: true,
		7: true,
		8: true,
	}
)

func ParseIntCode(inp string) []int {
	return utils.SplitAndParseInt(inp, ",")
}

func intSize(v int) int {
	return int(math.Log10(float64(v))) + 1
}

func parseOp(op int) []int {
	v := []int{}
	p := 100
	
	for op != 0 {
		part := op % p
		op -= part

		if p != 100 {
			part /= p/10
		}

		p *= 10
		v = append(v, part)
	}

	return v
}

type Computer struct {
	Input []int
	Output []int
	Code []int

	cur int
}

func (c *Computer) getParams(amt int, opData []int) []int {
	data := []int{}

	for i := 0; i < amt; i++ {
		m := 0
		if i + 1 < len(opData) {
			m = opData[i + 1]
		}
		if write_ops[opData[0]] && i == amt - 1 {
			m = 1
		}

		d := c.Code[c.cur + i + 1]

		switch m {
		case 0:
			data = append(data, c.Code[d])
		case 1:
			data = append(data, d)
		}
	}

	return data
}

func (c *Computer) doOp() {
	opData := parseOp(c.Code[c.cur])
	paramCount := param_amounts[opData[0]]

	params := c.getParams(paramCount, opData)

	switch opData[0] {
	case 1:
		c.Code[params[2]] = params[0] + params[1]
	case 2:
		c.Code[params[2]] = params[0] * params[1]
	case 3:
		c.Code[params[0]] = c.Input[0]
	case 4:
		c.Output = append(c.Output, params[0])
	case 5:
		if params[0] != 0 {
			c.cur = params[1]
			return
		}
	case 6:
		if params[0] == 0 {
			c.cur = params[1]
			return
		}
	case 7:
		v := 0
		if params[0] < params[1] {
			v = 1
		}
		c.Code[params[2]] = v
	case 8:
		v := 0
		if params[0] == params[1] {
			v = 1
		}
		c.Code[params[2]] = v
	case 99:
		return
	default:
		panic("Unknown op! " + strconv.Itoa(opData[0]))
	}

	c.cur += paramCount + 1
}

func (c *Computer) RunIntCode() {
	for c.Code[c.cur] != 99 {
		c.doOp()
	}
}
