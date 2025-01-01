package day9

import (
	"slices"

	"github.com/shadiestgoat/aoc/utils/sparse"
)

func ValidateInput(_pre []int, v int) bool {
	pre := make([]int, len(_pre))
	copy(pre, _pre)

	slices.Sort(pre)

	for i := 0; i < len(pre); i++ {
		for j := 0; j < len(pre); j++ {
			if i == j {
				continue
			}
			if pre[i]+pre[j] == v {
				return true
			}
		}
	}

	return false
}

func mostGenericSolver1(inp []int, size int) int {
	for i, v := range inp[size:] {
		if !ValidateInput(inp[i:i+size], v) {
			return v
		}
	}

	return -1
}

func GenericSolver1(inp string, size int) int {
	return mostGenericSolver1(sparse.SplitAndParseInt(inp, "\n"), size)
}

func Solve1(inp string) any {
	return GenericSolver1(inp, 25)
}

func GenericSolver2(inp []int, badNum int) (int, int) {
	minI := 0
	maxI := 1

	sum := inp[0] + inp[1]

	for {
		switch {
		case sum == badNum:
			return minI, maxI
		case sum > badNum:
			sum -= inp[minI]
			minI++
		case sum < badNum:
			maxI++
			sum += inp[maxI]
		}
	}
}

func Solve2(inp string) any {
	parsed := sparse.SplitAndParseInt(inp, "\n")

	badNum := mostGenericSolver1(parsed, 25)

	min, max := GenericSolver2(parsed, badNum)

	// Fun Fact: we no longer care about the parsed data
	arr := parsed[min : max+1]
	slices.Sort(arr)

	return arr[0] + arr[len(arr)-1]
}
