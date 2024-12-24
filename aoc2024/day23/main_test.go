package day23_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2024/day23"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day23.Solve1, 7)
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day23.Solve2, `co,de,ka,ta`)
}
