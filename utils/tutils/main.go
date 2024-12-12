package tutils

import (
	"strconv"
	"strings"
	"testing"
)

// Asserts e == r, and if not fails the test (FailNow)
func Assert(t *testing.T, exp, real any) {
	if exp == real {
		return
	}

	t.Fatalf("Grrr expected %v, but got %v", exp, real)
}

func AssertFunc(t *testing.T, inp string, f func(string) any, exp any) {
	r := f(strings.TrimSpace(inp))

	Assert(t, exp, r)
}

func AssertMany(t *testing.T, tests [][2]any, f func (string) any) {
	for i, cfg := range tests {
		t.Run(strconv.Itoa(i + 1), func(t *testing.T) {
			AssertFunc(t, cfg[0].(string), f, cfg[1])
		})
	}
}