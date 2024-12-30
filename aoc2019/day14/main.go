package day14

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

// !! wow thats a lot
const PART_2_ORE = 1_000_000_000_000

type Cache struct {
	Count int
	BankLefts map[string]int
}

type State struct {
	// res -> Rule
	rules map[string]*Rule
}

type Rule struct {
	Req

	Reqs []Req
}

type Req struct {
	Elm string
	Count int
}

func parseReq(s string) Req {
	spl := strings.Split(s, " ")
	return Req{
		Elm:   spl[1],
		Count: utils.ParseInt(spl[0]),
	}
}

func parseInput(inp string) *State {
	s := &State{
		rules:    map[string]*Rule{},
	}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l, " => ")
		r := &Rule{
			Req:  parseReq(spl[1]),
			Reqs: utils.SplitAndParseFunc(spl[0], ", ", parseReq),
		}

		s.rules[r.Elm] = r
	}

	return s
}

func (s State) oreAmount(fuelAmount int) int {
	queue := []Req{
		{"FUEL", fuelAmount},
	}

	ore := 0
	bank := map[string]int{}

	for len(queue) != 0 {
		cur := queue[len(queue) - 1]
		if cur.Elm == "ORE" {
			queue = queue[:len(queue) - 1]
			ore += cur.Count
			continue
		}

		banked := bank[cur.Elm]
		if banked != 0 {
			if cur.Count >= banked {
				cur.Count -= banked
				delete(bank, cur.Elm)
			} else {
				bank[cur.Elm] -= cur.Count
				cur.Count = 0
			}
		}

		queue = queue[:len(queue) - 1]
		if cur.Count == 0 {
			continue
		}

		r := s.rules[cur.Elm]
		mul := cur.Count/r.Count

		if cur.Count % r.Count != 0 {
			mul++
			bank[cur.Elm] = (mul * r.Count) - cur.Count
		}

		queue = append(queue, utils.Map(r.Reqs, func(r Req) Req { r.Count *= mul; return r })...)
	}

	return ore
}

func Solve1(inp string) any {
	s := parseInput(inp)

	return s.oreAmount(1)
}

func Solve2(inp string) any {
	s := parseInput(inp)

	forOneFuel := s.oreAmount(1)

	minBoundary := PART_2_ORE/forOneFuel
	maxBoundary := minBoundary + minBoundary/4

	found := 0
	for {
		if minBoundary == maxBoundary {
			found = minBoundary
			break
		}
		if maxBoundary - minBoundary == 1 {
			found = minBoundary
			break
		}

		mid := minBoundary + (maxBoundary - minBoundary)/2

		curOre := s.oreAmount(mid)
		if curOre == PART_2_ORE {
			found = mid
			break
		}

		if curOre > PART_2_ORE {
			maxBoundary = mid
		} else {
			minBoundary = mid
		}
	}

	return found
}
