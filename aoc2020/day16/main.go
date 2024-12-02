package day16

import (
	"strings"

	"shadygoat.eu/aoc/utils"
)

type Rule struct {
	Name string
	Rules []*Constraint
}

func (r Rule) Check(v int) bool {
	for _, c := range r.Rules {
		if c.Check(v) {
			return true
		}
	}

	return false
}

type Constraint struct {
	Min, Max int
}

func (c Constraint) Check(v int) bool {
	return c.Min <= v && v <= c.Max
}

func parseInput(inp string) ([]*Rule, []int, [][]int) {
	parts := strings.Split(inp, "\n\n")
	rules := []*Rule{}

	for _, l := range strings.Split(parts[0], "\n") {
		r := &Rule{}

		spl := strings.Split(l, ": ")
		
		r.Name = spl[0]

		ranges := strings.Split(spl[1], " or ")

		for _, ra := range ranges {
			c := strings.Split(ra, "-")

			r.Rules = append(r.Rules, &Constraint{
				Min: utils.ParseInt(c[0]),
				Max: utils.ParseInt(c[1]),
			})
		}

		rules = append(rules, r)
	}

	spl := strings.Split(parts[1], "\n")
	myTicket := utils.SplitAndParseInt(spl[1], ",")

	spl = strings.Split(parts[2], "\n")
	nearby := make([][]int, len(spl) - 1)

	for i, v := range spl[1:] {
		nearby[i] = utils.SplitAndParseInt(v, ",")
	}

	return rules, myTicket, nearby
}

func ValidateTicket(rules []*Rule, vals []int, onBad func (v int) bool) bool {
	ok := true

	for _, v := range vals {
		valid := false

		for _, r := range rules {
			if r.Check(v) {
				valid = true
				break
			}
		}

		if !valid {
			if onBad(v) {
				return false
			}

			ok = false
		}
	}

	return ok
}

func Solve1(inp string) any {
	rules, _, nearby := parseInput(inp)

	tot := 0

	for _, t := range nearby {
		ValidateTicket(rules, t, func(v int) bool {
			tot += v

			return false
		})
	}

	return tot
}

func Solve2(inp string) any {
	rules, own, nearby := parseInput(inp)

	goodNearby := [][]int{}

	for _, t := range nearby {
		if ValidateTicket(rules, t, func(v int) bool {return true}) {
			goodNearby = append(goodNearby, t)
		}
	}

	// Each index here is the same as a ticket value index. The array is rule indexes
	candidates := make([]map[int]bool, len(own))

	// Whenever a value is determined, we place it here
	determined := make([]int, len(own))
	// Rule index -> field index
	determinedRules := map[int]int{}

	fieldVals := make([][]int, len(own))

	for _, t := range append(goodNearby, own) {
		for i, v := range t {
			fieldVals[i] = append(fieldVals[i], v)
		}
	}

	for fi, field := range fieldVals {
		applicableRules := []int{}

		for ri, r := range rules {
			applicable := true

			for _, v := range field {
				if !r.Check(v) {
					applicable = false
					break
				}
			}

			if applicable {
				applicableRules = append(applicableRules, ri)
			}
		}

		candidates[fi] = map[int]bool{}

		for _, r := range applicableRules {
			candidates[fi][r] = true
		}

		if len(applicableRules) == 1 {
			r := applicableRules[0]

			determinedRules[r] = fi
			determined[fi] = r
		}
	}

	lastLen := len(determinedRules)
	for len(determinedRules) != len(rules) {
		for fi, cand := range candidates {
			for r := range cand {
				if _, ok := determinedRules[r]; ok {
					delete(cand, r)
				}
			}

			if len(cand) == 1 {
				r := 0

				// Ssshh jank code is fun
				for _r := range cand {
					r = _r
					break
				}

				determined[fi] = r
				determinedRules[r] = fi
			}
		}

		if lastLen == len(determinedRules) {
			panic("Making no progress")
		}
	}

	t := 1
	for i, r := range rules {
		if strings.HasPrefix(r.Name, "departure ") {
			fi := determinedRules[i]
			t *= own[fi]
		}
	}

	return t
}
