package day13

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/xy"
)

type Claw struct {
	// Strength of buttons A & B
	A, B xy.XY

	PriceLoc xy.XY
}

func (c Claw) Confirm(a, b int) bool {
	return c.PriceLoc == xy.XY{c.A[0]*a + c.B[0]*b, c.A[1]*a + c.B[1]*b}
}

func (c Claw) MinBtnPresses() (int, int) {
	// Its basically simultaneous equations
	// x = Ax * a + Bx * b
	// y = Ay * a + By * b
	//
	// Ax * a = x - Bx * b
	// a = (x - Bx * b)/Ax
	//
	// y = Ay * (x - Bx * b)/Ax + By * b
	// y * Ax = Ay * x - Ay * Bx * b + By * Ax * b
	// y * Ax - Ay * x = b * (By * Ax - Ay * Bx)
	//
	// b = (y * Ax - x * Ay)/(By * Ax - Ay * Bx)

	b := (c.PriceLoc[1]*c.A[0] - c.PriceLoc[0]*c.A[1]) / (c.B[1]*c.A[0] - c.A[1]*c.B[0])
	a := (c.PriceLoc[0] - c.B[0]*b) / c.A[0]

	return a, b
}

func parseInput(inp string, add int) []*Claw {
	claws := []*Claw{}

	for _, mac := range strings.Split(inp, "\n\n") {
		lines := strings.Split(mac, "\n")

		c := &Claw{
			A:        xy.XYFromArr(sparse.SplitAndParseInt(lines[0][11:], ", Y")),
			B:        xy.XYFromArr(sparse.SplitAndParseInt(lines[1][11:], ", Y")),
			PriceLoc: xy.XYFromArr(sparse.SplitAndParseInt(lines[2][9:], ", Y=")).Add(xy.XY{add, add}),
		}

		claws = append(claws, c)
	}

	return claws
}

func Solve1(inp string) any {
	claws := parseInput(inp, 0)

	tot := 0
	for _, c := range claws {
		a, b := c.MinBtnPresses()
		if a > 100 || b > 100 || a < 0 || b < 0 {
			continue
		}
		if !c.Confirm(a, b) {
			continue
		}

		tot += a*3 + b
	}

	return tot
}

func Solve2(inp string) any {
	claws := parseInput(inp, 10000000000000)

	tot := 0
	for _, c := range claws {
		a, b := c.MinBtnPresses()
		if a < 0 || b < 0 {
			continue
		}
		if !c.Confirm(a, b) {
			continue
		}

		tot += a*3 + b
	}

	return tot
}
