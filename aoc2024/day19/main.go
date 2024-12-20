package day19

import (
	"strings"
)

type RuleSet struct {
	Patterns []string
	Cache map[string]int
}

func parseInput(inp string) (*RuleSet, []string) {
	spl := strings.Split(inp, "\n\n")

	r := &RuleSet{
		Patterns:     strings.Split(spl[0], ", "),
		Cache:        map[string]int{},
	}

	return r, strings.Split(spl[1], "\n")
}

func (r *RuleSet) DesignPossibilityCounter(cur string, exitEarly bool) int {
	if v, ok := r.Cache[cur]; ok {
		return v
	}
	if cur == "" {
		return 1
	}

	t := 0
	for _, p := range r.Patterns {
		if !strings.HasPrefix(cur, p) {
			continue
		}

		poss := r.DesignPossibilityCounter(cur[len(p):], exitEarly)
		t += poss

		if poss != 0 && exitEarly {
			break
		}
	}

	r.Cache[cur] = t
	return t
}

// Hacky and slow? Yes.
func Solve1(inp string) any {
	r, designs := parseInput(inp)

	t := 0
	for _, d := range designs {
		t += r.DesignPossibilityCounter(d, true)
	}

	return t
}

func Solve2(inp string) any {
	r, designs := parseInput(inp)

	t := 0
	for _, d := range designs {
		t += r.DesignPossibilityCounter(d, false)
	}

	return t
}
