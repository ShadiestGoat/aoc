package day19

import (
	"fmt"
	"strings"

	"shadygoat.eu/aoc/utils"
)

type Rule struct {
	Possibilities map[string]bool

	Deps map[int]bool
	Resolved bool
	RawRule string
}

func resolveRule(rules map[int]*Rule, current []string, restOfRules []int) []string {
	if len(restOfRules) == 0 {
		return current
	}

	newPoss := []string{}

	for _, c := range current {
		for p := range rules[restOfRules[0]].Possibilities {
			newPoss = append(newPoss, c + p)
		}
	}

	return resolveRule(rules, newPoss, restOfRules[1:])
}

func parseAllRules(inp string) map[int]*Rule {
	rules := map[int]*Rule{}

	resolvedCount := 0

	for _, l := range strings.Split(inp, "\n") {
		r := &Rule{
			Deps: map[int]bool{},
		}

		spl := strings.Split(l, ": ")
		ruleID := utils.ParseInt(spl[0])

		r.RawRule = spl[1]

		if spl[1][0] == '"' {
			r.Possibilities = map[string]bool{
				string(spl[1][1]): true,
			}
			r.Resolved = true
			resolvedCount++
		} else {
			tmp := strings.ReplaceAll(spl[1], " | ", " ")

			for _, v := range utils.SplitAndParseInt(tmp, " ") {
				r.Deps[v] = true
			}
		}

		rules[ruleID] = r
	}

	lastResolvedCount := 0

	for lastResolvedCount != resolvedCount {
		fmt.Println(resolvedCount, len(rules))
		lastResolvedCount = resolvedCount

		for _, r := range rules {
			if r.Resolved {
				continue
			}
	
			resolvable := true
			for d := range r.Deps {
				if !rules[d].Resolved {
					resolvable = false
					break
				}
			}
	
			if !resolvable {
				continue
			}
	
			pipes := strings.Split(r.RawRule, " | ")
			resolved := map[string]bool{}

			for _, p := range pipes {
				givenRules := utils.SplitAndParseInt(p, " ")

				tmp := resolveRule(rules, []string{""}, givenRules)
				for _, v := range tmp {
					resolved[v] = true
				}
			}
	
			r.Resolved = true
			r.Possibilities = resolved
			resolvedCount++
		}
	}

	return rules
}

func Solve1(inp string) any {
	parts := strings.Split(inp, "\n\n")

	rules := parseAllRules(parts[0])
	fmt.Println("parsed")

	r0 := rules[0]

	count := 0
	for _, l := range strings.Split(parts[1], "\n") {
		if r0.Possibilities[l] {
			count++
		}
	}

	return count
}

func Solve2(inp string) any {
	return nil
}
