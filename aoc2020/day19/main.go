package day19

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Rule struct {
	ID int
	Val rune
	Definition [][]int
}

type RuleSet struct {
	Rules map[int]*Rule
}

// If returns true, then do not proceed
func (r *RuleSet) resolve(id int, cur []string, target string) ([]string, bool) {
	tr := r.Rules[id]

	if tr.Val != 0 {
		if len(cur) == 0 {
			if target[0] != byte(tr.Val) {
				return nil, true
			} else {
				return []string{string(tr.Val)}, false
			}
		}

		newCur := []string{}
		for _, c := range cur {
			v := c + string(tr.Val)
			if strings.HasPrefix(target, v) {
				newCur = append(newCur, v)
			}
		}

		return newCur, len(newCur) == 0
	}

	v := map[string]bool{}
	doResolve := false

	for _, def := range tr.Definition {
		newCurs := slices.Clone(cur)
		noResolve := false

		for _, rID := range def {
			nc, nr := r.resolve(rID, newCurs, target)
			if nr {
				noResolve = nr
				break
			}

			newCurs = nc
		}

		if noResolve {
			continue
		}

		doResolve = true
		utils.MapListKeysOnExisting(newCurs, v)
	}

	if !doResolve {
		return nil, true
	}

	return utils.MapKeys(v), false
}

func (r *RuleSet) IsGood(s string) bool {
	possibilities, noResolve := r.resolve(0, []string{}, s)
	if noResolve {
		return false
	}

	for _, p := range possibilities {
		if p == s {
			return true
		}
	}

	return false
}

// Override: map[RuleID] -> the rule
func parseInput(inp string, override map[int]string) (*RuleSet, []string) {
	spl := strings.Split(inp, "\n\n")

	rs := &RuleSet{
		Rules: map[int]*Rule{},
	}

	for _, v := range strings.Split(spl[0], "\n") {
		exprSpl := strings.Split(v, ": ")
		id := utils.ParseInt(exprSpl[0])

		expr := exprSpl[1]
		if override[id] != "" {
			expr = override[id]
		}

		r := &Rule{
			ID:    id,
			Val:   0,
			Definition: [][]int{},
		}
		if expr[0] == '"' {
			r.Val = rune(expr[1])
		} else {
			r.Definition = utils.SplitAndParseInt2(expr, " | ", " ")
		}

		rs.Rules[id] = r
	}

	return rs, strings.Split(spl[1], "\n")
}

func Solve1(inp string) any {
	rs, checks := parseInput(inp, map[int]string{})
	count := 0

	for _, c := range checks {
		if rs.IsGood(c) {
			count++
		}
	}

	return count
}

func Solve2(inp string) any {
	// Fuck you, creator, it can handle any possibilities due to efficient prefix matching!!
	rs, checks := parseInput(inp, map[int]string{
		8: `42 | 42 8`,
		11: `42 31 | 42 11 31`,
	})
	count := 0

	for _, c := range checks {
		if rs.IsGood(c) {
			count++
		}
	}

	return count
}
