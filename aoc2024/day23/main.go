package day23

import (
	"math/bits"
	"slices"
	"strings"
)

type State struct {
	// ID -> (Conn ID -> true)
	Conns map[uint64]*StupidSet

	IDMap map[uint64]string
	IDRevMap map[string]uint64
	lastID uint64

	longest StupidSet
	tested map[StupidSet]bool
}

// I thought bit operations would be fun...
type StupidSet [10]uint64
const SET_LEN uint64 = 10 * 64

func (s StupidSet) And(s2 StupidSet) StupidSet {
	res := StupidSet{}

	for i, p := range s {
		p2 := s2[i]
		res[i] = p & p2
	}

	return res
}

func (s StupidSet) Copy() StupidSet {
	return s
}

func (s *StupidSet) Set(id uint64) {
	s[id/64] |= 1 << (id % 64)
}

func (s StupidSet) At(id uint64) bool {
	return ((s[id/64] >> (id % 64)) & 1) == 1
}

func (s StupidSet) Len() int {
	t := 0
	for _, p := range s {
		t += bits.OnesCount64(p)
	}

	return t
}

func (s *State) getID(v string) uint64 {
	if id, ok := s.IDRevMap[v]; ok {
		return id
	}

	s.lastID++
	s.IDRevMap[v] = s.lastID
	s.IDMap[s.lastID] = v

	return s.lastID
}

func (s *State) addConn(a, b string) {
	aID, bID := s.getID(a), s.getID(b)

	if s.Conns[aID] == nil {
		se := &StupidSet{}
		s.Conns[aID] = se
	}
	
	s.Conns[aID].Set(bID)
}

// Stateful recursion >:3 (evil)
func (s *State) recLanFinder(past StupidSet, t uint64) {
	if past.And(*s.Conns[t]).Len() < past.Len() {
		return
	}

	cur := past.Copy()
	cur.Set(t)

	if s.tested[cur] {
		return
	}
	s.tested[cur] = true

	if cur.Len() > s.longest.Len() {
		s.longest = cur
	}

	for id := range s.Conns {
		if cur.At(id) {
			continue
		}

		s.recLanFinder(cur, id)
	}
}

func parseInput(inp string) *State {
	s := &State{
		Conns:    map[uint64]*StupidSet{},
		IDMap:    map[uint64]string{},
		IDRevMap: map[string]uint64{},
		lastID:   0,
		longest:  StupidSet{},
		tested:   map[StupidSet]bool{},
	}

	for _, v := range strings.Split(inp, "\n") {
		spl := strings.Split(v, "-")
		s.addConn(spl[0], spl[1])
		s.addConn(spl[1], spl[0])
	}

	return s
}

func Solve1(inp string) any {
	s := parseInput(inp)

	found := map[StupidSet]bool{}

	for a, conns := range s.Conns {
		if s.IDMap[a][0] != 't' {
			continue
		}
		
		for b := uint64(0); b < SET_LEN; b++ {
			if a == b || !conns.At(b) {
				continue
			}

			common := conns.And(*s.Conns[b])
			if common.Len() == 0 {
				continue
			}

			for c := uint64(0); c < SET_LEN; c++ {
				if c == a || c == b || !common.At(c) {
					continue
				}
				
				s := StupidSet{}
				s.Set(a)
				s.Set(b)
				s.Set(c)

				found[s] = true
			}
		}
	}

	return len(found)
}

func Solve2(inp string) any {
	s := parseInput(inp)

	for a := range s.Conns {
		s.recLanFinder(StupidSet{}, a)
	}

	longestChain := []string{}
	for i := uint64(0); i < SET_LEN; i++ {
		if s.longest.At(i) {
			longestChain = append(longestChain, s.IDMap[i])
		}
	}

	slices.Sort(longestChain)

	return strings.Join(longestChain, ",")
}
