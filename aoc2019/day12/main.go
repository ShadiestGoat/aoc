package day12

import (
	"fmt"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type World struct {
	Moons []*Moon
}

func (w World) String() string {
	str := []string{}

	for _, m := range w.Moons {
		str = append(str, m.String())
	}

	return strings.Join(str, "\n")
}

type Moon struct {
	Vel [3]int
	Pos [3]int
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.Pos[0], m.Pos[1], m.Pos[2], m.Vel[0], m.Vel[1], m.Vel[2])
}

func (m *Moon) Move() {
	for i, v := range m.Vel {
		m.Pos[i] += v
	}
}

func (w *World) Tick() {
	for i, m := range w.Moons {
		for _, m2 := range w.Moons[i + 1:] {
			for j, p := range m.Pos {
				if p > m2.Pos[j] {
					m2.Vel[j]++
					m.Vel[j]--
				} else if p < m2.Pos[j] {
					m2.Vel[j]--
					m.Vel[j]++
				}
			}
		}
	}

	for _, m := range w.Moons {
		m.Move()
	}
}

func parseInput(inp string) *World {
	w := &World{
		Moons: []*Moon{},
	}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l[1:len(l)-1], ", ")
		m := &Moon{
			Vel: [3]int{},
			Pos: [3]int{},
		}

		for i, p := range spl {
			m.Pos[i] = utils.ParseInt(p[2:])
		}

		w.Moons = append(w.Moons, m)
	}

	return w
}

func GenericSolve1(inp string, steps int) any {
	w := parseInput(inp)

	for i := 0; i < steps; i++ {
		w.Tick()
	}

	t := 0
	for _, m := range w.Moons {
		t += utils.AbsSum(m.Pos[:]) * utils.AbsSum(m.Vel[:])
	}

	return t
}

func Solve1(inp string) any {
	return GenericSolve1(inp, 1000)
}

func Solve2(inp string) any {
	w := parseInput(inp)
	initPos := [3][4]int{}

	for i, m := range w.Moons {
		for j, p := range m.Pos {
			initPos[j][i] = p
		}
	}

	axisPeriods := [3][2]int{}

	i := 0
	for {
		w.Tick()

		goodPos := [3]bool{true, true, true}
		for mID, m := range w.Moons {
			for j := 0; j < 3; j++ {
				if m.Vel[j] != 0 || m.Pos[j] != initPos[j][mID] {
					goodPos[j] = false
				}
			}
		}

		for j, g := range goodPos {
			if !g {
				continue
			}

			for k, v := range axisPeriods[j] {
				if v == 0 {
					axisPeriods[j][k] = i
					break
				}
			}
		}

		allPeriodsFound := true
		for _, p := range axisPeriods {
			for _, v := range p {
				if v == 0 {
					allPeriodsFound = false
					break
				}
			}

			if !allPeriodsFound {
				break
			}
		}

		if allPeriodsFound {
			break
		}

		i++
	}

	periods := [3]int{}

	for i := 0; i < 3; i++ {
		periods[i] = axisPeriods[i][1] - axisPeriods[i][0]
	}

	return utils.LCM(periods[:]...)
}
