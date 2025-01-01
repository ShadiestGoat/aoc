package day25

import "github.com/shadiestgoat/aoc/utils/sparse"

func parseInput(inp string) []int {
	return sparse.SplitAndParseInt(inp, "\n")
}

func transformOnce(v int, subject int) int {
	return (v * subject) % 20201227
}

func findLoopSize(publicKey int) int {
	i := 0

	cur := 1
	for {
		cur = transformOnce(cur, 7)
		if cur == publicKey {
			return i + 1
		}

		i++
	}
}

func Solve1(inp string) any {
	pubs := parseInput(inp)

	loopSize := findLoopSize(pubs[0])

	cur := 1
	for i := 0; i < loopSize; i++ {
		cur = transformOnce(cur, pubs[1])
	}

	return cur
}

func Solve2(inp string) any {
	return nil
}
