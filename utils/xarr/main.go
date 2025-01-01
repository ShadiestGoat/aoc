package xarr

import "strings"

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
