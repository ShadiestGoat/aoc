package day6

import (
	"slices"
	"strings"
)

type Orbit struct {
	// Own name
	Name string

	// Direct orbit object
	Direct string
	// Anything else that this this thing orbits around
	Indirect []string
}

type World struct {
	// Center of Mass -> Things that orbit around it
	Orbits map[string][]*Orbit

	OrbitNames map[string]*Orbit
}

func (w *World) resolveOrbit(curPath []string, o *Orbit) {
	o.Indirect = curPath[:len(curPath) - 1]
	for _, no := range w.Orbits[o.Name] {
		w.resolveOrbit(append(slices.Clone(curPath), o.Name), no)
	}
}

func parseInput(inp string) *World {
	w := &World{
		Orbits: map[string][]*Orbit{},
		OrbitNames: map[string]*Orbit{},
	}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l, ")")

		o := &Orbit{
			Name:   spl[1],
			Direct: spl[0],
		}

		w.OrbitNames[spl[1]] = o
		w.Orbits[spl[0]] = append(w.Orbits[spl[0]], o)
	}

	for _, o := range w.Orbits["COM"] {
		w.resolveOrbit([]string{"COM"}, o)
	}

	return w
}

func Solve1(inp string) any {
	w := parseInput(inp)
	c := 0

	for _, o := range w.OrbitNames {
		// +1 for direct
		c += len(o.Indirect) + 1
	}

	return c
}

func Solve2(inp string) any {
	w := parseInput(inp)

	myDeps := w.OrbitNames["YOU"].Indirect
	sanDeps := w.OrbitNames["SAN"].Indirect
	distance := 0

	for i := len(sanDeps) - 1; i >= 0; i-- {
		idx := slices.Index(myDeps, sanDeps[i])
		if idx == -1 {
			continue
		}

		distance = len(myDeps) - idx + len(sanDeps) - i

		break
	}

	return distance
}
