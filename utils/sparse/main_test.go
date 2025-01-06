package sparse_test

import (
	"strconv"
	"testing"

	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func TestParseInt(t *testing.T) {
	cases := []string{
		"-5",
		"+5",
		"5",
		"125",
		"999",
		"-555",
		"+555",
	}

	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			exp, _ := strconv.Atoi(c)

			tutils.AssertFunc(t, c, sparse.ParseInt, exp)
		})
	}
}