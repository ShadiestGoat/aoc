package tutils

import (
	"strconv"
	"strings"
	"testing"
)

func AssertCustomCompare[T any](t *testing.T, exp, real T, cmp func (a, b T) bool) {
	if cmp(exp, real) {
		return
	}

	failAssert(t, exp, real)
}

func failAssert(t *testing.T, exp, real any) {
	t.Fatalf("Grrr expected %v, but got %v", exp, real)
}

// Asserts e == r, and if not fails the test (FailNow)
func Assert(t *testing.T, exp, real any) {
	if exp == real {
		return
	}
	failAssert(t, exp, real)
}

func AssertFunc(t *testing.T, inp string, f func(string) any, exp any) {
	r := f(strings.TrimSpace(inp))

	Assert(t, exp, r)
}

func AssertFuncCustomCompare[T any](t *testing.T, inp string, f func(string) T, exp T, cmp func (a, b T) bool) {
	r := f(strings.TrimSpace(inp))

	AssertCustomCompare(t, exp, r, cmp)
}

func AssertMany(t *testing.T, tests [][2]any, f func (string) any) {
	for i, cfg := range tests {
		t.Run(strconv.Itoa(i + 1), func(t *testing.T) {
			t.Log(cfg[0])

			AssertFunc(t, cfg[0].(string), f, cfg[1])
		})
	}
}