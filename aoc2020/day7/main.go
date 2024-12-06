package day7

import (
	"strconv"
	"strings"
)

type Bag struct {
	Name       string
	DirectDeps []*BagDep
	// Bag name -> bag amount, including deps of deps
	ResolvedDeps map[string]int
}

type BagDep struct {
	Count int
	Name  string
}

func parseBag(bag string) *Bag {
	d := strings.SplitN(bag, " bags contain ", 2)
	if len(d) != 2 {
		panic("Bad parse of " + bag)
	}
	if d[1] == "no other bags." {
		return &Bag{
			Name:       d[0],
			DirectDeps: []*BagDep{},
		}
	}

	deps := []*BagDep{}
	rawDeps := strings.Split(d[1][:len(d[1])-1], ", ")

	for _, rawDep := range rawDeps {
		v := strings.SplitN(rawDep, " ", 4)

		// Error handling is for nerds
		c, _ := strconv.Atoi(v[0])

		deps = append(deps, &BagDep{
			Count: c,
			Name:  v[1] + " " + v[2],
		})
	}

	return &Bag{
		Name:         d[0],
		DirectDeps:   deps,
		ResolvedDeps: nil,
	}
}

func resolveDep(all map[string]*Bag, deps []*BagDep) map[string]int {
	cur := map[string]int{}

	for _, d := range deps {
		cur[d.Name] += d.Count

		if b, ok := all[d.Name]; !ok {
			cur[d.Name] += d.Count
		} else {
			if b.ResolvedDeps == nil {
				resp := resolveDep(all, b.DirectDeps)
				b.ResolvedDeps = resp
			}

			for n, c := range b.ResolvedDeps {
				cur[n] += c * d.Count
			}
		}
	}

	return cur
}

func ParseInput(inp string) map[string]*Bag {
	m := map[string]*Bag{}

	for _, d := range strings.Split(inp, "\n") {
		bag := parseBag(d)
		m[bag.Name] = bag
	}

	for _, b := range m {
		resp := resolveDep(m, b.DirectDeps)
		b.ResolvedDeps = resp
	}

	return m
}

func Solve1(inp string) any {
	o := ParseInput(inp)
	tot := 0

	for _, v := range o {
		if v.ResolvedDeps["shiny gold"] >= 1 {
			tot++
		}
	}

	return tot
}

func Solve2(inp string) any {
	o := ParseInput(inp)
	tot := 0

	for _, v := range o["shiny gold"].ResolvedDeps {
		tot += v
	}

	return tot
}
