package day21

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils/xy"
	"github.com/shadiestgoat/aoc/utils/sparse"
)

var (
	keyPadCoords = map[rune]xy.XY{
		'7': {-2, -3},
		'8': {-1, -3},
		'9': {0, -3},

		'4': {-2, -2},
		'5': {-1, -2},
		'6': {0, -2},

		'1': {-2, -1},
		'2': {-1, -1},
		'3': {0, -1},

		'0': {-1, 0},
		'A': {0, 0},
	}

	dirKeyPadCoords = map[xy.XY]xy.XY{
		xy.DIR_UP: {-1, 0},
		{}:           {0, 0},

		xy.DIR_LEFT:  {-2, 1},
		xy.DIR_DOWN:  {-1, 1},
		xy.DIR_RIGHT: {0, 1},
	}
)

type Dir xy.XY

func (d Dir) String() string {
	switch d {
	case Dir{}:
		return "A"
	case Dir(xy.DIR_UP):
		return "^"
	case Dir(xy.DIR_DOWN):
		return "v"
	case Dir(xy.DIR_LEFT):
		return "<"
	case Dir(xy.DIR_RIGHT):
		return ">"
	}

	return "?"
}

func diffToDirSeq(diff xy.XY, og xy.XY) []Dir {
	base := diff.Unit()
	seq := []Dir{}

	doAction := func(i int) {
		for j := 0; j != diff[i]; j += base[i] {
			c := Dir{}
			c[i] = base[i]

			seq = append(seq, c)
		}
	}

	yFirst := func() {
		doAction(1)
		doAction(0)
	}
	xFirst := func() {
		doAction(0)
		doAction(1)
	}

	if og[1]+diff[1] == 0 && og[0] == -2 {
		xFirst()
	} else if og[0]+diff[0] == -2 && og[1] == 0 {
		yFirst()
	} else if diff[0] < 0 {
		xFirst()
	} else {
		yFirst()
	}

	t := og
	for _, v := range seq {
		t = t.Add(xy.XY(v))
		if t == (xy.XY{-2, 0}) {
			panic(">:(")
		}
	}

	return seq
}

// [oldPos, diff]
type Cache = map[[2]xy.XY]map[int]int

func diffDeepResolve(lastPos, newPos xy.XY, left int, cache Cache) (int, xy.XY) {
	diff := newPos.Add(lastPos.Mul(-1))

	cacheKey := [2]xy.XY{lastPos, diff}
	if steps, ok := cache[cacheKey]; ok {
		if steps[left] != 0 {
			return steps[left], newPos
		}
	} else {
		cache[cacheKey] = map[int]int{}
	}

	seq := append(diffToDirSeq(diff, lastPos), Dir{})
	if left == 0 {
		return len(seq), newPos
	}

	seqLen := 0
	recPos := dirKeyPadCoords[xy.XY{}]
	for _, v := range seq {
		s, np := diffDeepResolve(recPos, dirKeyPadCoords[xy.XY(v)], left-1, cache)
		seqLen += s
		recPos = np
	}

	cache[cacheKey][left] = seqLen

	return seqLen, newPos
}

func genericSolve(inp string, robotCount int) int {
	lines := strings.Split(inp, "\n")
	amt := 0

	var cache Cache = Cache{}

	for _, code := range lines {
		lastCode := 'A'
		l := 0

		for _, c := range code {
			sl, _ := diffDeepResolve(keyPadCoords[lastCode], keyPadCoords[c], robotCount, cache)
			lastCode = c
			l += sl
		}

		num := sparse.ParseInt(code[:len(code)-1])
		amt += l * num
	}

	return amt
}

func Solve1(inp string) any {
	return genericSolve(inp, 2)
}

func Solve2(inp string) any {
	return genericSolve(inp, 25)
}
