package day11_test

import (
	"strings"
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day11"
)

const INPUT = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

const STATE_P1_R1 = `
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`

const STATE_P1_R2 = `
#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
`

const STATE_P1_R3 = `
#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##
`

const STATE_P2_R1 = `
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`

const STATE_P2_R2 = `
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
`

const STATE_P2_R3 = `
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
`

func TestStateExec(t *testing.T) {
	newState := func() *day11.State {
		return day11.ParseInput(strings.TrimSpace(INPUT))
	}

	assertState := func(s *day11.State, t *testing.T, i int, exp string) {
		if s.CurState != strings.TrimSpace(exp) {
			t.Log("-- expected --\n" + exp)
			t.Log("-- gotten --\n\n" + s.CurState)
			t.Errorf("-- States mismatch (#%v) --", i+1)
		}
	}

	t.Run("vis=1", func(t *testing.T) {
		allExp := []string{STATE_P1_R1, STATE_P1_R2, STATE_P1_R3}
		s := newState()

		for i, exp := range allExp {
			s.Exec(1, 4)
			assertState(s, t, i, exp)
		}
	})

	t.Run("vis=Inf", func(t *testing.T) {
		allExp := []string{STATE_P2_R1, STATE_P2_R2, STATE_P2_R3}
		s := newState()

		for i, exp := range allExp {
			s.Exec(9999, 5)
			assertState(s, t, i, exp)
		}
	})
}

func TestSolve2(t *testing.T) {
	resp := day11.Solve2(strings.TrimSpace(INPUT))

	if resp != 26 {
		t.Fatalf(":((((( Expected 26, got %v", resp)
	}
}
