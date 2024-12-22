package day22

import (
	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) []int {
	return utils.SplitAndParseInt(inp, "\n")
}

func MixAndPrune(a, b int) int {
	return (a ^ b) % 16777216
}

func EvolveSecret(a int) int {
	v := MixAndPrune(a, a * 64)
	v = MixAndPrune(v, v/32)
	return MixAndPrune(v, v * 2048)
}

type Ring [4]int
func (r *Ring) Add(n int) {
	for i := 0; i < 3; i++ {
		r[i] = r[i + 1]
	}
	r[3] = n
}

func Solve1(inp string) any {
	nums := parseInput(inp)

	t := 0
	for _, v := range nums {
		for i := 0; i < 2000; i++ {
			v = EvolveSecret(v)
		}
		
		t += v
	}

	return t
}

func Solve2(inp string) any {
	nums := parseInput(inp)

	seqTests := map[Ring]int{}
	for _, v := range nums {
		lastPrice := v % 10
		curSeq := Ring{}
		tested := map[Ring]bool{}

		for i := 0; i < 2000; i++ {
			v = EvolveSecret(v)
			curPrice := v % 10

			curSeq.Add(curPrice - lastPrice)
			lastPrice = curPrice
			if i < 3 || tested[curSeq] {
				continue
			}

			tested[curSeq] = true
			seqTests[curSeq] += curPrice
		}
	}

	best := 0
	for _, v := range seqTests {
		if v > best {
			best = v
		}
	}

	return best
}
