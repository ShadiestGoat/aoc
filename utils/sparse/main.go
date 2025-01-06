package sparse

import (
	"strconv"
	"strings"
)

func SplitAndParseInt(inp string, delim string) []int {
	return SplitAndParseFunc(inp, delim, ParseInt)
}

// Its useful shut up
func SplitAndParseInt2(inp string, d1, d2 string) [][]int {
	return SplitAndParseFunc(inp, d1, func(s string) []int {
		return SplitAndParseInt(s, d2)
	})
}

func SplitAndParseFunc[T any](inp string, delim string, h func(s string) T) []T {
	segs := strings.Split(inp, delim)
	o := make([]T, len(segs))

	for i, s := range segs {
		o[i] = h(s)
	}

	return o
}

func SplitParseAndScan[T any](inp string, delim string, p func (s string) T, h func (T)) {
	for _, v := range strings.Split(inp, delim) {
		h(p(v))
	}
}

func ParseInt(v string) int {
	n, _ := strconv.Atoi(v)
	return n
	// m := 1
	// if v[0] == '-' || v[0] == '+' {
	// 	if v[0] == '-' {
	// 		m = -1
	// 	}
	// 	v = v[1:]
	// }

	// n := 0
	// for _, r := range v {
	// 	if r < '0' || r > '9' {
	// 		panic("Bad number syntax: " + v)
	// 	}

	// 	n = n*10 + int(r - '0')
	// }

	// return n * m
}
