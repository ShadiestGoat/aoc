package day20

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

func coolMatchLine(str string, match string) bool {
	for i, r := range match {
		if r == ' ' {
			continue
		}

		if r != rune(str[i]) {
			return false
		}
	}

	return true
}

type CoolMatcher struct {
	Coords []utils.XY
	maxX int
}

func (m *CoolMatcher) matchSpecific(offset utils.XY, str []string) bool {
	for _, rc := range m.Coords {
		c := rc.Add(offset)
		if c.OutOfBounds(utils.XY{len(str[0]) - 1, len(str) - 1}) {
			return false
		}

		if str[c[1]][c[0]] != '#' {
			return false
		}
	}

	return true
}

func (m *CoolMatcher) Match(str string) map[utils.XY]bool {
	matched := map[utils.XY]bool{}
	lines := strings.Split(str, "\n")
	maxOffset := len(lines[0]) - m.maxX - 1

	for i := range lines {
		for o := 0; o <= maxOffset; o++ {
			off := utils.XY{o, i}
			if !m.matchSpecific(off, lines) {
				continue
			}

			for _, c := range m.Coords {
				matched[c.Add(off)] = true
			}
		}
	}

	return matched
}

func NewCoolMatcher(str string, rawMatcher string) *CoolMatcher {
	matcher := []utils.XY{}
	maxMatcherX := 0
	for i, v := range strings.Split(strings.TrimSpace(rawMatcher), "\n") {
		ml := v[1:len(v) - 1]
		maxX := len(ml) - 1
		if maxX > maxMatcherX {
			maxMatcherX = maxX
		}

		for j, r := range ml {
			if r == ' ' {
				continue
			}

			matcher = append(matcher, utils.XY{j, i})
		}
	}

	return &CoolMatcher{
		Coords: matcher,
		maxX:   maxMatcherX,
	}
}
