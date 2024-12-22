package day21

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

var (
	keyPadCoords = map[rune]utils.XY{
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

	dirKeyPadCoords = map[utils.XY]utils.XY{
		utils.DIR_UP: {-1, 0},
		{}:           {0, 0},

		utils.DIR_LEFT:  {-2, 1},
		utils.DIR_DOWN:  {-1, 1},
		utils.DIR_RIGHT: {0, 1},
	}
)

type Dir utils.XY

func (d Dir) String() string {
	switch d {
	case Dir{}:
		return "A"
	case Dir(utils.DIR_UP):
		return "^"
	case Dir(utils.DIR_DOWN):
		return "v"
	case Dir(utils.DIR_LEFT):
		return "<"
	case Dir(utils.DIR_RIGHT):
		return ">"
	}

	return "?"
}

func diffToDirSeq(diff utils.XY, og utils.XY) []Dir {
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

	if og[1] + diff[1] == 0 && og[0] == -2 {
		xFirst()
	} else if og[0] + diff[0] == -2 && og[1] == 0 {
		yFirst()
	} else if diff[0] < 0 {
		xFirst()
	} else {
		yFirst()
	}

	t := og
	for _, v := range seq {
		t = t.Add(utils.XY(v))
		if t == (utils.XY{-2, 0}) {
			panic(">:(")
		}
	}

	return seq
}

// [oldPos, diff]
var cache = map[[2]utils.XY]map[int]int{}

func diffDeepResolve(lastPos, newPos utils.XY, left int) (int, utils.XY) {
	diff := newPos.Add(lastPos.Mul(-1))

	cacheKey := [2]utils.XY{lastPos, diff}
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
	recPos := dirKeyPadCoords[utils.XY{}]
	for _, v := range seq {
		s, np := diffDeepResolve(recPos, dirKeyPadCoords[utils.XY(v)], left - 1)
		seqLen += s
		recPos = np
	}

	cache[cacheKey][left] = seqLen

	return seqLen, newPos
}

func genericSolve(inp string, robotCount int) int {
	lines := strings.Split(inp, "\n")
	amt := 0

	for _, code := range lines {
		lastCode := 'A'
		l := 0

		for _, c := range code {
			sl, _ := diffDeepResolve(keyPadCoords[lastCode], keyPadCoords[c], robotCount)
			lastCode = c
			l += sl
		}
		
		num := utils.ParseInt(code[:len(code)-1])
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
