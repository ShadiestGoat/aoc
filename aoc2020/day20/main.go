package day20

import (
	"fmt"
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Tile struct {
	ID int
	Data []string
}

type BorderInfo struct {
	BoardID int
	// 0 - top, 3 - left
	BorderPos int
	BorderFlipped bool
}

func parseInput(inp string) (map[string][]*BorderInfo, map[int]*Tile) {
	idToTile := map[int]*Tile{}
	borders := map[string][]*BorderInfo{}

	addBorder := func (id, pos int, border string) {
		borders[border] = append(borders[border], &BorderInfo{
			BoardID:       id,
			BorderPos:     pos,
			BorderFlipped: false,
		})

		rb := []rune(border)
		slices.Reverse(rb)

		borders[string(rb)] = append(borders[string(rb)], &BorderInfo{
			BoardID:       id,
			BorderPos:     pos,
			BorderFlipped: true,
		})
	}

	for _, raw := range strings.Split(inp, "\n\n") {
		spl := strings.SplitN(raw, "\n", 2)

		id := utils.ParseInt(spl[0][5:len(spl[0]) - 1])
		data := strings.Split(spl[1], "\n")

		addBorder(id, 0, data[0])
		addBorder(id, 2, data[len(data) - 1])

		lb, rb := make([]rune, len(data)), make([]rune, len(data))
		for i, r := range data {
			lb[i] = rune(r[0])
			rb[i] = rune(r[len(r) - 1])
		} 

		addBorder(id, 1, string(lb))
		addBorder(id, 3, string(rb))

		idToTile[id] = &Tile{id, data}
	}

	return borders, idToTile
}

func matchCount(borders map[string][]*BorderInfo) map[int]int {
	idToMatchCount := map[int]int{}
	for _, info := range borders {
		if len(info) == 1 {
			continue
		}

		for _, b := range info {
			idToMatchCount[b.BoardID]++
		}
	}

	return idToMatchCount
}

// God this is so clever -- why bother doing all the border stuff when you could find the # of matches
func Solve1(inp string) any {
	borders, _ := parseInput(inp)

	idToMatchCount := matchCount(borders)

	tot := 1
	for id, v := range idToMatchCount {
		if v == 4 {
			tot *= id
		}
	}

	return tot
}

type BuildState struct {
	Built [][]string
	Borders map[string][]*BorderInfo
	Tiles map[int]*Tile
	Corners map[int]bool
}

func (b *BuildState) getBorder(x, y int, pos int) string {
	d := b.Built[y][x]
	lines := strings.Split(d, "\n")

	switch pos {
	case 0:
		return lines[0]
	case 2:
		return lines[len(lines) - 1]
	case 1, 3:
		v := make([]rune, len(lines))
		for i, l := range lines {
			r := l[0]
			if pos == 3 {
				r = l[len(l) - 1]
			}

			v[i] = rune(r)
		}
	}

	return ""
}

func (b *BuildState) build() (string, bool) {
	var (
		y = 0
		x = 1
		cornerCount = 0
	)

	for {
		bo := b.getBorder(x - 1, 0, 1)
		info := b.Borders[bo]
		if len(info) == 1 {
			return "", false
		}
		
	}
}

// Assume no false matches....... (except double flips)
func makeBoard(borders map[string][]*BorderInfo, tiles map[int]*Tile) string {
	// matches := matchCount(borders)

	// tID, flip 
	choice := [3]int{}
	for id, c := range matches {
		if c != 4 {
			continue
		}

		break
	}

	for _, info := range borders {
		if len(info) == 1 {
			fmt.Println(info[0])
			continue
		}
		fmt.Println(info[0], info[1])
	}

	return ""
}

func Solve2(inp string) any {
	board := makeBoard(parseInput(inp))

	return board
}
