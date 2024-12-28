package utils

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

func ParseInt(v string) int {
	p, err := strconv.Atoi(v)
	PanicIfErr(err, "parsing '%v' as int", v)

	return p
}

func MapLines(inp string, h func(string)) {
	for _, l := range strings.Split(inp, "\n") {
		h(l)
	}
}

func MapListKeys[T comparable](l []T) map[T]bool {
	m := make(map[T]bool, len(l))

	MapListKeysOnExisting(l, m)

	return m
}

func MapListKeysOnExisting[T comparable](l []T, m map[T]bool) {
	for _, v := range l {
		m[v] = true
	}
}

func MapKeys[T comparable, V any](m map[T]V) []T {
	v := make([]T, 0, len(m))
	for k := range m {
		v = append(v, k)
	}

	return v
}

func Map[T any, V any](s []T, h func (T) V) []V {
	nv := make([]V, len(s))

	for i, v := range s {
		nv[i] = h(v)
	}

	return nv
}

func Reduce[T any, R any](initial R, arr []T, f func (acc R, cur T) R) R {
	acc := initial

	for _, v := range arr {
		acc = f(acc, v)
	}

	return acc
}

func Sum(arr []int) int {
	return Reduce(0, arr, func(acc int, cur int) int {
		return acc + cur
	})
}

func AbsSum(arr []int) int {
	return Reduce(0, arr, func(acc, cur int) int {
		if cur < 0 {
			cur = -cur
		}

		return acc + cur
	})
}
