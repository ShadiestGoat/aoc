package day17

import (
	"slices"
	"strconv"
	"strings"

	"github.com/shadiestgoat/aoc/utils/sparse"
	"github.com/shadiestgoat/aoc/utils/xarr"
)

type State struct {
	InsPtr  int
	AllCode []int
	A, B, C int
	Output  []int
}

func (s *State) String() string {
	str := []string{}

	for i, v := range []int{s.A, s.B, s.C} {
		str = append(str, "Registry "+string(rune('A'+i))+": "+strconv.Itoa(v))
	}

	return strings.Join(str, "\n")
}

func (s *State) adv(operand int) (int, bool) {
	cc := s.comboCode(operand) - 1
	if cc < 0 {
		return -1, false
	}

	return s.A / (2 << (s.comboCode(operand) - 1)), true
}

func (s *State) runOnce() bool {
	operand := s.AllCode[s.InsPtr+1]
	ok := true

	switch s.AllCode[s.InsPtr] {
	case 0:
		// Pow of 2 is super fun w/ ints -- god I love bitshifts
		s.A, ok = s.adv(operand)
	case 1:
		s.B ^= operand
	case 2:
		s.B = s.comboCode(operand) % 8
	case 3:
		if s.A == 0 {
			break
		}
		s.InsPtr = operand

		return true
	case 4:
		s.B ^= s.C
	case 5:
		s.Output = append(s.Output, s.comboCode(operand)%8)
	case 6:
		s.B, ok = s.adv(operand)
	case 7:
		s.C, ok = s.adv(operand)
	}

	s.InsPtr += 2
	return ok
}

func (s *State) comboCode(c int) int {
	switch c {
	case 0, 1, 2, 3:
		return c
	case 4:
		return s.A
	case 5:
		return s.B
	case 6:
		return s.C
	}

	panic("Unknown combo code: " + strconv.Itoa(c))
}

func (s *State) runForever() {
	for s.InsPtr < len(s.AllCode) {
		s.runOnce()
	}
}

func parseInput(inp string) *State {
	spl := strings.Split(inp, "\n\n")

	s := &State{
		AllCode: sparse.SplitAndParseInt(spl[1][9:], ","),
	}

	for _, v := range strings.Split(spl[0], "\n") {
		rowData := strings.Split(v, ": ")
		v := sparse.ParseInt(rowData[1])

		switch rowData[0][len(rowData[0])-1] {
		case 'A':
			s.A = v
		case 'B':
			s.B = v
		case 'C':
			s.C = v
		}
	}

	return s
}

func Solve1(inp string) any {
	s := parseInput(inp)
	s.runForever()

	return strings.Join(xarr.Map(s.Output, strconv.Itoa), ",")
}

func Solve2(inp string) any {
	s := parseInput(inp)

	operant := -1
	for i := 0; i < len(s.AllCode); i += 2 {
		if s.AllCode[i] == 0 {
			if operant != -1 {
				panic("Got bad assumption: multiple A writes (opcode 0)")
			}

			operant = s.AllCode[i+1]
		}
	}

	if operant > 3 {
		panic("Bad Assumption: non-static division")
	}

	p2 := 2 << (operant - 1)

	base := 0
	tI := len(s.AllCode) - 2

	for tI >= 0 {
		found := false

		for i := 0; i < 64; i++ {
			s.A, s.B, s.C = base+i, 0, 0
			s.InsPtr = 0
			s.Output = []int{}

			s.runForever()

			if slices.Equal(s.Output, s.AllCode[tI:]) {
				base = (base + i) * p2
				tI--

				found = true

				break
			}
		}

		if !found {
			panic("Didn't find the thing :(")
		}
	}

	return base / 8
}
