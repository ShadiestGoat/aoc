package day21_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day21"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day21.Solve1, 5)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day21.Solve2, `mxmxvkd,sqjhc,fvjkl`)
}