package day7

import (
	"fmt"
	"maps"
	"slices"

	"github.com/shadiestgoat/aoc/aoc2019/intcode"
)

func permutations[T any](sl []T, cur int) [][]T {
	if cur == len(sl) {
		return [][]T{slices.Clone(sl)}
	}

	o := [][]T{}

	for i := cur; i < len(sl); i++ {
		sl[i], sl[cur] = sl[cur], sl[i]
		o = append(o, permutations(sl, cur+1)...)
		// Go back <3
		sl[i], sl[cur] = sl[cur], sl[i]
	}

	return o
}

func Solve1(inp string) any {
	code := intcode.ParseIntCode(inp)
	allPerms := permutations([]int{4, 3, 2, 1, 0}, 0)
	bestV := 0

	for _, curSetup := range allPerms {
		lastO := 0

		for _, phase := range curSetup {
			lastO = intcode.QuickRun(code, []int{phase, lastO})[0]
		}

		if bestV == 0 || lastO > bestV {
			bestV = lastO
		}
	}

	return bestV
}

func Solve2(inp string) any {
	code := intcode.ParseIntCode(inp)
	allPerms := permutations([]int{5, 6, 7, 8, 9}, 0)
	bestV := 0

	for _, curSetup := range allPerms {
		lastO := 0
		allAmps := []*intcode.Computer{}

		for _, phase := range curSetup {
			allAmps = append(allAmps, &intcode.Computer{
				Input:  []int{phase},
				Code:   maps.Clone(code),
			})
		}

		for {
			finHalt := false

			for i, amp := range allAmps {
				amp.Input = append(amp.Input, lastO)
				haltReason := amp.RunIntCode()

				lastO = amp.Output[0]
				if len(amp.Output) > 1 {
					fmt.Println("Ai")
				}

				amp.Output = []int{}

				if haltReason && i == len(allAmps) - 1 {
					finHalt = true
					break
				}
			}

			if finHalt {
				break
			}
		}

		if bestV == 0 || lastO > bestV {
			bestV = lastO
		}
	}

	return bestV
}
