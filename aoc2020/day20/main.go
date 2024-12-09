package day20

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

const (
	PART_2_MATCH = `
|                  # |
|#    ##    ##    ###|
| #  #  #  #  #  #   |
`
)

type Tile struct {
	ID int
	Data []string
}

type BorderInfo struct {
	TileID int
	// 0 - top, 3 - left
	BorderPos int
	BorderFlipped bool
}

func parseInput(inp string) (map[string][]*BorderInfo, map[int]*Tile) {
	idToTile := map[int]*Tile{}
	borders := map[string][]*BorderInfo{}

	addBorder := func (id, pos int, border string) {
		borders[border] = append(borders[border], &BorderInfo{
			TileID:       id,
			BorderPos:     pos,
			BorderFlipped: false,
		})

		rb := []rune(border)
		slices.Reverse(rb)

		borders[string(rb)] = append(borders[string(rb)], &BorderInfo{
			TileID:       id,
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

		addBorder(id, 3, string(lb))
		addBorder(id, 1, string(rb))

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
			idToMatchCount[b.TileID]++
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

func Solve2(inp string) any {
	board := makeBoard(parseInput(inp))
	if board == "" {
		panic("Empty Board!")
	}

	tot := 0

	for _, l := range strings.Split(board, "\n") {
		for _, r := range l {
			if r == '#' {
				tot++
			}
		}
	}

	matcher := NewCoolMatcher(inp, PART_2_MATCH)

	lines := strings.Split(board, "\n")
	for rot := 0; rot < 4; rot++ {
		rotLines := rotateString(lines, rot)

		for _, fx := range []bool{true, false} {
			for _, fy := range []bool{true, false} {
				lines := rotLines
				if fx {
					lines = flipX(lines)
				}
				if fy {
					lines = flipY(lines)
				}

				matched := matcher.Match(strings.Join(lines, "\n"))
				if len(matched) != 0 {
					return tot - len(matched)
				}
			}
		}
	}

	return nil
}
