package day2_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/day2"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
1,9,10,3,2,3,11,0,99,30,40,50
`

func TestRunIntCode(t *testing.T) {
	var testCases = [][2]string{
		{`1,0,0,0,99`, `2,0,0,0,99`},
		{`2,3,0,3,99`, `2,3,0,6,99`},
		{`2,4,4,5,99,0`, `2,4,4,5,99,9801`},
		{`1,1,1,4,99,5,6,0,99`, `30,1,1,4,2,5,6,0,99`},		
	}

	for _, cfg := range testCases {
		t.Run(cfg[0], func(t *testing.T) {
			code := day2.ParseInput(cfg[0])
			day2.RunIntCode(code)

			resp := strings.Join(utils.Map(code, strconv.Itoa), ",")
			tutils.Assert(t, cfg[1], resp)
		})
	}
}

// func TestSolve2(t *testing.T) {
//	tutils.AssertFunc(t, INPUT, day2.Solve2, VALUE)
//}
