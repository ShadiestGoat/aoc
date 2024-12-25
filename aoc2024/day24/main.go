package day24

import (
	"fmt"
	"maps"
	// "maps"
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type GateOp string

const (
	OP_AND GateOp = "AND"
	OP_OR GateOp = "OR"
	OP_XOR GateOp = "XOR"
)

func (g GateOp) Do(a, b bool) bool {
	switch g {
	case OP_AND:
		return a && b
	case OP_OR:
		return a || b
	case OP_XOR:
		return a != b
	}

	return false
}

type Gate struct {
	Op GateOp
	Nums []string
	Res  string
}

type State struct {
	Nums map[string]bool
	Gates []*Gate
	OutToGate map[string]*Gate
	swaps []string
}

func (s State) getNum(prefix rune) int {
	o := 0
	for k, v := range s.Nums {
		if k[0] != byte(prefix) || !v {
			continue
		}

		p := utils.ParseInt(k[1:])
		o |= 1 << p
	}

	return o
}

func (s *State) Do(op GateOp, res string, a, b bool) {
	if res[0] == 'x' || res[0] == 'y' {
		panic("????")
	}
	s.Nums[res] = op.Do(a, b)
}

func (s *State) DoAll() bool {
	done := map[int]bool{}
	lastLen := 0

	for {
		for i, g := range s.Gates {
			if done[i] {
				continue
			}

			a, ok := s.Nums[g.Nums[0]]
			if !ok {
				continue
			}

			b, ok := s.Nums[g.Nums[1]]
			if !ok {
				continue
			}

			s.Do(g.Op, g.Res, a, b)
			done[i] = true
		}

		if len(done) == len(s.Gates) {
			return true
		}
		if len(done) == lastLen {
			return false
		}
		lastLen = len(done)
	}
}

func (s *State) find(op GateOp, n1, n2 string) *Gate {
	for _, g := range s.Gates {
		if g.Op != op {
			continue
		}
		if slices.Contains(g.Nums, n1) && slices.Contains(g.Nums, n2) {
			return g
		}
	}

	return nil
}

func parseInput(inp string) *State {
	spl := strings.Split(inp, "\n\n")

	s := &State{
		Nums:  map[string]bool{},
		Gates: []*Gate{},
		OutToGate: map[string]*Gate{},
	}

	for _, l := range strings.Split(spl[1], "\n") {
		words := strings.Split(l, " ")

		g := &Gate{
			Op:   GateOp(words[1]),
			Nums: []string{words[0], words[2]},
			Res:  words[4],
		}

		s.Gates = append(s.Gates, g)
		s.OutToGate[g.Res] = g
	}

	for _, l := range strings.Split(spl[0], "\n") {
		parts := strings.Split(l, ": ")
		s.Nums[parts[0]] = parts[1] == "1"
	}

	// Simple sort isn't possible, despite my tries :(

	return s
}

func Solve1(inp string) any {
	s := parseInput(inp)
	s.DoAll()

	return s.getNum('z')
}

func fmtN(s rune, i int) string {
	return string(s) + fmt.Sprintf("%02d", i)
}

func (s *State) swapOut(a, b string) {
	s.OutToGate[a].Res = b
	s.OutToGate[b].Res = a

	s.OutToGate[a], s.OutToGate[b] = s.OutToGate[b], s.OutToGate[a]
	s.swaps = append(s.swaps, a, b)
}

func (s *State) fixI(i int) bool {
	curOut := s.OutToGate[fmtN('z', i)]
	
	lastNums := s.OutToGate[fmtN('z', i - 1)].Nums

	// Must feed into real out
	xyXor := s.find(OP_XOR, fmtN('x', i), fmtN('y', i))

	carry := s.find(OP_AND, lastNums[0], lastNums[1])
	xy1And := s.find(OP_AND, fmtN('x', i - 1), fmtN('y', i - 1))
	lastOr := s.find(OP_OR, carry.Res, xy1And.Res)
	realOut := s.find(OP_XOR, lastOr.Res, xyXor.Res)

	// Sanity check...
	if curOut == realOut {
		return false
	}

	if realOut == nil {
		if curOut.Op != OP_XOR {
			fmt.Println("Swaps", s.swaps)
			panic("Grr")
		}

		if slices.Contains(curOut.Nums, lastOr.Res) {
			n := 0
			if curOut.Nums[0] == lastOr.Res {
				n = 1
			}

			s.swapOut(xyXor.Res, curOut.Nums[n])
		} else {
			n := 0
			if curOut.Nums[0] == xyXor.Res {
				n = 1
			}

			s.swapOut(lastOr.Res, curOut.Nums[n])
		}
	} else {
		s.swapOut(fmtN('z', i), realOut.Res)
	}

	
	return true
}

func Solve2(inp string) any {
	s := parseInput(inp)
	baseNums := maps.Clone(s.Nums)

	zCount := 0
	for _, g := range s.Gates {
		if g.Res[0] == 'z' {
			zCount++
		}
	}

	for {
		s.DoAll()
		res, z := s.getNum('x') + s.getNum('y'), s.getNum('z')
		if res == z {
			break
		}
		s.Nums = maps.Clone(baseNums)

		fixed := false
		for i := 0; i < zCount; i++ {			
			if i > 3 && i != zCount - 1 && s.fixI(i) {
				fixed = true
				break
			}
		}
		
		if !fixed {
			fmt.Println(s.swaps)
			panic("No fixes...")
		}
	}

	slices.Sort(s.swaps)

	return strings.Join(s.swaps, ",")
}
