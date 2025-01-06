package funiter_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/utils/funiter"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func TestScan(t *testing.T) {
	tests := [][2]string{
		{"123@321@456", "@"},
		{"123@321@456@", "@"},
		{"123@321@@456", "@"},
		{"@123@321@@456@", "@"},
		{"@@123@321@@456@@", "@"},
		{"123@@321@@", "@@"},
		{"123@@321@@@321", "@@"},
		{"@@@@321@@@321", "@@"},
		{"@@@@321@@@321", "@@@"},
	}

	for _, cfg := range tests {
		t.Run(cfg[0], func(t *testing.T) {
			exp := strings.Split(cfg[0], cfg[1])
			real := []string{}

			funiter.Scan(cfg[0], cfg[1], func(s string) bool {
				real = append(real, s)
				return false
			})

			tutils.AssertSlices(t, exp, real)
		})
	}
}
