package day11_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day11"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
125 17
`

func testGameRunner[T any](t *testing.T, parse func (string) T, h func (T, int) int) {
	tutils.AssertFunc(t, INPUT, func(s string) any {
		v := parse(s)

		return h(v, 25)
	}, 55312)
}

func TestRunGameList(t *testing.T) {
	testGameRunner(t, day11.ParseInputList, day11.RunGameList)
}

func TestRunGameArray(t *testing.T) {
	testGameRunner(t, day11.ParseInputArray, day11.RunGameArray)
}

func TestRunGameState(t *testing.T) {
	testGameRunner(t, day11.ParseInputArray, day11.RunGameState)
}

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day11.Solve1, 55312)
}