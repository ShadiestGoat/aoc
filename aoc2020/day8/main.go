package day8

import (
	"strconv"
	"strings"
)

type CurState struct {
	Acc             int
	Index           int
	LastIndex       int
	History         map[int]bool
	AllInstructions []string
}

func (s *CurState) Reset() {
	s.Acc = 0
	s.Index = 0
	s.LastIndex = 0
	s.History = map[int]bool{}
}

func (s *CurState) Exec() {
	s.History[s.Index] = true

	cur := s.AllInstructions[s.Index]

	n, _ := strconv.Atoi(cur[4:])

	s.LastIndex = s.Index
	s.Index++

	switch cur[:3] {
	case "acc":
		s.Acc += n
	case "nop":
		return
	case "jmp":
		s.Index += n - 1
	}
}

// Execs until reaches end or history repeats itself
// Returns true if history is repeated
func (s *CurState) RepeatExec() bool {
	for {
		s.Exec()

		if s.History[s.Index] {
			return true
		}

		if len(s.AllInstructions) <= s.Index {
			return false
		}
	}
}

func ParseInput(inp string) *CurState {
	return &CurState{
		Acc:             0,
		Index:           0,
		LastIndex:       0,
		History:         map[int]bool{},
		AllInstructions: strings.Split(inp, "\n"),
	}
}

func Solve1(inp string) any {
	s := ParseInput(inp)

	s.RepeatExec()

	return s.Acc
}

func Solve2(inp string) any {
	s := ParseInput(inp)

	s.RepeatExec()

	curCandidates := []int{}

	for h := range s.History {
		if s.AllInstructions[h][:3] == "acc" {
			continue
		}

		curCandidates = append(curCandidates, h)
	}

	switchIns := func(i int) {
		oldIns := s.AllInstructions[i]
		newIns := ""

		switch oldIns[:3] {
		case "acc":
			panic("Somehow last instruction was acc")
		case "jmp":
			newIns = "nop"
		case "nop":
			newIns = "jmp"
		}

		s.AllInstructions[i] = newIns + oldIns[3:]
	}

	for {
		switchIns(curCandidates[0])

		if !s.RepeatExec() {
			return s.Acc
		}

		switchIns(curCandidates[0])

		curCandidates = curCandidates[1:]

		s.Reset()
	}
}
