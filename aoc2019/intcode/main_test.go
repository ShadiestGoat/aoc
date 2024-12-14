package intcode_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2019/intcode"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func TestRunIntCode(t *testing.T) {
	// these tests assume a limited memory!
	var testCases = [][2]any{
		{`1,0,0,0,99`, `2,0,0,0,99`},
		{`2,3,0,3,99`, `2,3,0,6,99`},
		{`2,4,4,5,99,0`, `2,4,4,5,99,9801`},
		{`1,1,1,4,99,5,6,0,99`, `30,1,1,4,2,5,6,0,99`},
		{`1002,4,3,4,33`, `1002,4,3,4,99`},
	}

	tutils.AssertMany(t, testCases, func(s string) any {
		comp := &intcode.Computer{
			Code: intcode.ParseIntCode(s),
		}
		comp.RunIntCode()

		arr := make([]int, len(comp.Code))
		for i, v := range comp.Code {
			arr[i] = v
		}

		return strings.Join(utils.Map(arr, strconv.Itoa), ",")
	})

	var ioTestCases = [][3]any{
		{`3,9,8,9,10,9,4,9,99,-1,8`, 8, 1},
		{`3,9,8,9,10,9,4,9,99,-1,8`, 11, 0},
		{`3,3,1108,-1,8,3,4,3,99`, 8, 1},
		{`3,3,1108,-1,8,3,4,3,99`, 11, 0},
		{`3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9`, 0, 0},
		{`3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9`, 999, 1},
		{`3,3,1105,-1,9,1101,0,0,12,4,12,99,1`, 0, 0},
		{`3,3,1105,-1,9,1101,0,0,12,4,12,99,1`, 999, 1},
		{`104,1125899906842624,99`, 0, 1125899906842624},
	}

	for _, cfg := range ioTestCases {
		t.Run(cfg[0].(string) + "-" + strconv.Itoa(cfg[1].(int)), func(t *testing.T) {
			comp := &intcode.Computer{
				Input: []int{cfg[1].(int)},
				Code:  intcode.ParseIntCode(cfg[0].(string)),
			}
			comp.RunIntCode()

			tutils.Assert(t, cfg[2], comp.Output[0])
		})
	}
}
