package day10

import (
	"slices"

	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/shadiestgoat/aoc/utils/sparse"
)

func Solve1(inp string) any {
	parsed := sparse.SplitAndParseInt(inp, "\n")

	slices.Sort(parsed)

	diffs := map[int]int{
		3: 1, // 1 is for the final -> device jump
	}

	for i, v := range parsed {
		diff := 0

		if i == 0 {
			diff = v
		} else {
			diff = v - parsed[i-1]
		}

		diffs[diff]++
	}

	return diffs[1] * diffs[3]
}

type Group struct {
	Numbers []int
	Solid   bool
}

func (g Group) LastN() int {
	return g.Numbers[len(g.Numbers)-1]
}

func PossibilityCounter(min, max int, path []int) int {
	c := 0

	// This is the one place I'll cave and use an external lib :(
	tmpPaths := append(
		combinations.All(path),
		[]int{},
	)

	for _, _tmp := range tmpPaths {
		tmp := append([]int{min}, append(_tmp, max)...)

		good := true
		for j, v := range tmp[1:] {
			if v-tmp[j] > 3 {
				good = false
				break
			}
		}

		if good {
			c++
		}
	}

	return c
}

func Solve2(inp string) any {
	parsed := sparse.SplitAndParseInt(inp, "\n")

	slices.Sort(parsed)

	groups := []*Group{
		{
			Numbers: []int{
				0,
			},
			Solid: true,
		},
	}

	for _, v := range parsed {
		lg := groups[len(groups)-1]

		maxDiff := v-lg.Numbers[len(lg.Numbers)-1] == 3

		if maxDiff {
			if lg.Solid {
				if len(lg.Numbers) == 2 {
					lg.Numbers[1] = v
				} else {
					lg.Numbers = append(lg.Numbers, v)
				}
			} else {
				lastN := lg.LastN()

				lg.Numbers = lg.Numbers[:len(lg.Numbers)-1]
				if len(lg.Numbers) == 0 {
					groups = groups[:len(groups)-1]
				}

				groups = append(groups, &Group{
					Numbers: []int{lastN, v},
					Solid:   true,
				})
			}
		} else {
			if lg.Solid {
				groups = append(groups, &Group{
					Numbers: []int{v},
					Solid:   false,
				})
			} else {
				lg.Numbers = append(lg.Numbers, v)
			}
		}
	}

	// not ideal but data merging is good

	groupsFixed := []*Group{
		groups[0],
	}

	for _, g := range groups[1:] {
		lg := groupsFixed[len(groupsFixed)-1]

		if lg.Solid && g.Solid {
			if len(lg.Numbers) == 1 {
				lg.Numbers = append(lg.Numbers, g.Numbers[1])
			} else {
				lg.Numbers[1] = g.Numbers[1]
			}
		} else {
			groupsFixed = append(groupsFixed, g)
		}
	}

	if !groupsFixed[len(groupsFixed)-1].Solid {
		lg := groupsFixed[len(groupsFixed)-1]
		ln := lg.LastN()

		lg.Numbers = lg.Numbers[:len(lg.Numbers)-1]
		groupsFixed = append(groupsFixed, &Group{
			Numbers: []int{ln},
			Solid:   true,
		})
	}

	groups = groupsFixed

	counter := 1

	for i, g := range groups {
		if g.Solid {
			continue
		}

		lg := groups[i-1]
		ng := groups[i+1]

		counter *= PossibilityCounter(lg.LastN(), ng.Numbers[0], g.Numbers)
	}

	return counter
}
