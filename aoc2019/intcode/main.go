package intcode

import (
	"math"
	"strconv"

	"github.com/shadiestgoat/aoc/utils/sparse"
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
		9: 1,
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

func ParseIntCode(inp string) map[int]int {
	m := map[int]int{}

	for i, v := range sparse.SplitAndParseInt(inp, ",") {
		m[i] = v
	}

	return m
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
	Code map[int]int

	cur int
	relativeBase int
}

func (c *Computer) getParams(amt int, opData []int) []int {
	data := []int{}

	for i := 0; i < amt; i++ {
		m := 0
		if i + 1 < len(opData) {
			m = opData[i + 1]
		}

		d := c.Code[c.cur + i + 1]

		if write_ops[opData[0]] && i == amt - 1 {
			if m == 2 {
				d += c.relativeBase
			}

			m = 1
		}


		switch m {
		case 0:
			data = append(data, c.Code[d])
		case 1:
			data = append(data, d)
		case 2:
			data = append(data, c.Code[d + c.relativeBase])
		}
	}

	return data
}

// Returns true if can continue
func (c *Computer) doOp() bool {
	opData := parseOp(c.Code[c.cur])
	paramCount := param_amounts[opData[0]]

	params := c.getParams(paramCount, opData)

	switch opData[0] {
	case 1:
		c.Code[params[2]] = params[0] + params[1]
	case 2:
		c.Code[params[2]] = params[0] * params[1]
	case 3:
		if len(c.Input) == 0 {
			return false
		}
		c.Code[params[0]] = c.Input[0]
		c.Input = c.Input[1:]
	case 4:
		c.Output = append(c.Output, params[0])
	case 5:
		if params[0] != 0 {
			c.cur = params[1]
			return true
		}
	case 6:
		if params[0] == 0 {
			c.cur = params[1]
			return true
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
	case 9:
		c.relativeBase += params[0]
	case 99:
		return false
	default:
		panic("Unknown op! " + strconv.Itoa(opData[0]))
	}

	c.cur += paramCount + 1

	return true
}

// Runs the code until either theres no more input, but input is needed OR until it halts
// If it halts, returns true
func (c *Computer) RunIntCode() bool {
	for c.doOp() {
	}

	return c.Code[c.cur] == 99
}

func (c *Computer) ConsumeOutput() []int {
	o := c.Output
	c.Output = []int{}
	return o
}

// Util function for cloning code & Quickly running the computer.
// The return value is the output
func QuickRun(code map[int]int, inp []int) []int {
	m := map[int]int{}
	for i, v := range code {
		m[i] = v
	}

	comp := &Computer{
		Input:  inp,
		Code:   m,
	}

	comp.RunIntCode()

	return comp.Output
}

func NewComp(inp string) *Computer {
	return &Computer{
		Input:        []int{},
		Output:       []int{},
		Code:         ParseIntCode(inp),
	}
}