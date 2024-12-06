package day17

import (
	"slices"
	"strings"
)

type Coords3 [3]int
type Coords4 [4]int

// Good code only
func AssignBS(c any, v []int) any {
	if c, ok := c.(Coords3); ok {
		copy(c[:], v)
		return c
	}

	if c, ok := c.(Coords4); ok {
		copy(c[:], v)
		return c
	}

	return nil
}

func GetBS(c any) []int {
	if c, ok := c.(Coords3); ok {
		return c[:]
	}

	if c, ok := c.(Coords4); ok {
		return c[:]
	}

	return nil
}

type CoordT interface {
	Coords3 | Coords4
}

type Map[CT CoordT] struct {
	Data map[CT]bool
}

func ParseInput[CT CoordT](inp string) *Map[CT] {
	m := &Map[CT]{
		Data: map[CT]bool{},
	}

	lines := strings.Split(inp, "\n")

	for i, l := range lines {
		for j, v := range l {
			if v != '#' {
				continue
			}

			var c CT

			c[0] = j
			c[1] = i

			m.Data[c] = true
		}
	}

	return m
}

func surroundingCoordsGeneric(og []int, vals [][]int, i int) [][]int {
	newVals := [][]int{}
	last := i == len(og)-1

	for _, v := range vals {
		noDiff := slices.Clone(v)
		noDiff[i] = og[i]

		p1 := slices.Clone(v)
		p1[i] = og[i] + 1

		m1 := slices.Clone(v)
		m1[i] = og[i] - 1

		newVals = append(newVals, [][]int{p1, m1}...)

		if !last || !slices.Equal(og, noDiff) {
			newVals = append(newVals, noDiff)
		}
	}

	if last {
		return newVals
	}

	return surroundingCoordsGeneric(og, newVals, i+1)
}

func SurroundingCoords(c []int) [][]int {
	return surroundingCoordsGeneric(c, [][]int{c}, 0)
}

func (m *Map[CT]) CountActive(coords []CT) int {
	count := 0

	for _, c := range coords {
		if m.Data[c] {
			count++
		}
	}

	return count
}

func (m Map[CT]) convertBS(og [][]int) []CT {
	goBS := make([]CT, len(og))

	for i, v := range og {
		var bs CT
		goBS[i] = AssignBS(bs, v).(CT)
	}

	return goBS
}

func (m *Map[CT]) CountActiveSurround(c CT) int {
	return m.CountActive(m.convertBS(SurroundingCoords(GetBS(c))))
}

func (m *Map[CT]) Exec() {
	n := map[CT]bool{}
	cachedSurrounds := map[CT]bool{}

	for c := range m.Data {
		surround := m.convertBS(SurroundingCoords(GetBS(c)))
		activeCount := m.CountActive(surround)

		if activeCount == 2 || activeCount == 3 {
			n[c] = true
		}

		for _, sc := range surround {
			if cachedSurrounds[sc] {
				continue
			}

			if m.CountActiveSurround(sc) == 3 {
				n[sc] = true
			}

			cachedSurrounds[sc] = true
		}
	}

	m.Data = n
}

func Solve1(inp string) any {
	m := ParseInput[Coords3](inp)

	for i := 0; i < 6; i++ {
		m.Exec()
	}

	return len(m.Data)
}

func Solve2(inp string) any {
	m := ParseInput[Coords4](inp)

	for i := 0; i < 6; i++ {
		m.Exec()
	}

	return len(m.Data)
}
