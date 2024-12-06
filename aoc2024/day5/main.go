package day5

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Config struct {
	// num -> it's deps (ie. stuff that has to be before it)
	NumDeps map[int][]int
	Lists   [][]int
}

func (c *Config) meetsDeps(n int, allVals map[int]bool, used map[int]bool) bool {
	for _, d := range c.NumDeps[n] {
		if !allVals[d] {
			continue
		}

		if !used[d] {
			return false
		}
	}

	return true
}

func (c *Config) isListRight(l []int) bool {
	used := map[int]bool{}
	allVals := utils.MapListKeys(l)

	for _, n := range l {
		if !c.meetsDeps(n, allVals, used) {
			return false
		}

		used[n] = true
	}

	return true
}

func (c *Config) makeListRight(l []int) []int {
	allVals := utils.MapListKeys(l)
	used := map[int]bool{}
	receptacle := []int{}
	newL := []int{}

	doReceptacle := func() {
		for {
			newRec := []int{}

			for _, r := range receptacle {
				if c.meetsDeps(r, allVals, used) {
					used[r] = true
					newL = append(newL, r)
					continue
				}

				newRec = append(newRec, r)
			}

			if len(newRec) == len(receptacle) {
				return
			}

			receptacle = newRec
		}
	}

	for _, n := range l {
		doReceptacle()

		if !c.meetsDeps(n, allVals, used) {
			receptacle = append(receptacle, n)
			continue
		}

		used[n] = true
		newL = append(newL, n)
	}

	doReceptacle()

	if len(receptacle) != 0 {
		panic(":/")
	}

	return newL
}

func parseInput(inp string) *Config {
	spl := strings.Split(inp, "\n\n")
	cfg := &Config{
		NumDeps: map[int][]int{},
		Lists:   utils.SplitAndParseInt2(spl[1], "\n", ","),
	}

	for _, l := range strings.Split(spl[0], "\n") {
		p := utils.SplitAndParseInt(l, "|")
		cfg.NumDeps[p[1]] = append(cfg.NumDeps[p[1]], p[0])
	}

	return cfg
}

func Solve1(inp string) any {
	cfg := parseInput(inp)
	sum := 0

	for _, l := range cfg.Lists {
		if !cfg.isListRight(l) {
			continue
		}

		sum += l[len(l)/2]
	}

	return sum
}

func Solve2(inp string) any {
	cfg := parseInput(inp)
	sum := 0

	for _, l := range cfg.Lists {
		if cfg.isListRight(l) {
			continue
		}

		nl := cfg.makeListRight(l)

		sum += nl[len(nl)/2]
	}

	return sum
}
