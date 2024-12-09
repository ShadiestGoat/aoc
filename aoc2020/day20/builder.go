package day20

import (
	"math"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

func (t *Tile) getBorder(pos int) string {
	switch pos {
	case 0:
		return t.Data[0]
	case 2:
		return t.Data[len(t.Data) - 1]
	case 1, 3:
		v := []rune{}
		for _, l := range t.Data {
			r := l[0]
			if pos == 1 {
				r = l[len(l) - 1]
			}

			v = append(v, rune(r))
		}

		return string(v)
	}

	return ""
}

func otherBorderPos(c int) int {
	// im tired shut up idc about doing normalization this is simpler
	switch c {
	case 0:
		return 2
	case 1:
		return 3
	case 2:
		return 0
	case 3:
		return 1
	}

	return 0
}


type BuildState struct {
	Built [][]*Tile
	Borders map[string][]*BorderInfo
	Tiles map[int]*Tile
	Corners map[int][]*BorderInfo
}

func (b *BuildState) getBorder(x, y int, pos int) string {
	d := b.Built[y][x]

	return d.getBorder(pos)
}

func (b *BuildState) drawDebug() string {
	str := ""

	for _, row := range b.Built {
		rowStr := make([]string, len(row[0].Data))

		for j, t := range row {
			for i, l := range t.Data {
				pfx := " "
				if j == 0 {
					pfx = ""
				}

				rowStr[i] += pfx + l
			}
		}

		str += "\n\n" + strings.Join(rowStr, "\n")
	}

	return str[2:]
}

func (b *BuildState) drawOutput() string {
	str := ""

	for _, row := range b.Built {
		rowStr := make([]string, len(row[0].Data))

		for _, t := range row {
			for i, l := range t.Data {
				rowStr[i] += l[1:len(l) - 1]
			}
		}

		str += "\n" + strings.Join(rowStr[1:len(rowStr) - 1], "\n")
	}

	return str[1:]
}

// Returns the tile t & the neighboring border pos
func (b *BuildState) getLastTile(curX, curY int) (t *Tile, tPos int) {
	tx, ty := curX - 1, curY
	tPos = 1

	if curX == 0 {
		tx, ty = 0, curY - 1
		tPos = 2
	}
	if ty >= len(b.Built) || tx >= len(b.Built[ty]) {
		return
	}

	t = b.Built[ty][tx]

	return
}

func (b *BuildState) build() (string, bool) {
	var (
		y = 0
		x = 1
		cornerCount = 1
	)

	newLine := func () {
		y++
		x = 0
		b.Built = append(b.Built, []*Tile{})
	}

	for {
		lastTile, lastTilePos := b.getLastTile(x, y)
		if lastTile == nil {
			return "", false
		}

		lastBorder := lastTile.getBorder(lastTilePos)
		info, ok := b.Borders[lastBorder]

		if !ok {
			utils.PrintJSON(lastBorder, lastTile, b.Borders)
			panic("Unknown border")
		}

		if len(info) == 1 {
			newLine()
			continue
		}

		var nextBorder *BorderInfo
		for _, v := range info {
			if v.TileID == lastTile.ID {
				continue
			}
			nextBorder = v
			break
		}

		nextID := nextBorder.TileID
		nextLines := b.Tiles[nextID].Data

		targetBorder := otherBorderPos(lastTilePos)
		rotationFlips := math.Abs(float64(nextBorder.BorderPos - targetBorder)) >= 2

		if nextBorder.BorderFlipped != rotationFlips {
			if nextBorder.BorderPos == 0 || nextBorder.BorderPos == 2 {
				nextLines = flipX(nextLines)
			} else {
				nextLines = flipY(nextLines)
			}
		}
	
		nextLines = rotateString(nextLines, targetBorder - nextBorder.BorderPos)

		b.Built[y] = append(b.Built[y], &Tile{
			ID:   nextID,
			Data: nextLines,
		})

		// Sanity check
		if y != 0 && x != 0 {
			if x >= len(b.Built[y - 1]) {
				return "", false
			}

			topBorder := b.getBorder(x, y - 1, 2)
			if b.getBorder(x, y, 0) != topBorder {
				if b.getBorder(x, y, 2) != topBorder {
					return "", false
				}

				b.Built[y][x].Data = flipY(b.Built[y][x].Data)
			}
		}

		x++
		if _, ok := b.Corners[nextID]; ok {
			cornerCount++
			if cornerCount == 4 {
				return b.drawOutput(), true
			}
		}
	}
}

func NewBuildState(cID int, fx, fy bool, corners map[int][]*BorderInfo, tiles map[int]*Tile, borders map[string][]*BorderInfo) *BuildState {
	topLeftNonMatches := [4]bool{}
	for _, v := range corners[cID] {
		topLeftNonMatches[v.BorderPos] = true
	}

	rot := 0
	for i, v := range append(topLeftNonMatches[1:], topLeftNonMatches[0]) {
		if topLeftNonMatches[i] && v {
			rot = i - 2
			break
		}
	}

	tl := tiles[cID].Data
	tl = rotateString(tl, rot)
	if fy {
		tl = flipY(tl)
	}
	if fx {
		tl = flipX(tl)
	}

	return &BuildState{
		Built:   [][]*Tile{
			{
				{
					ID:   cID,
					Data: tl,
				},
			},
		},
		Borders: borders,
		Tiles:   tiles,
		Corners: corners,
	}
}

// Assume no false matches....... (except double flips)
func makeBoard(borders map[string][]*BorderInfo, tiles map[int]*Tile) string {
	matches := matchCount(borders)
	
	corners := map[int][]*BorderInfo{}
	for tID, matchCount := range matches {
		if matchCount != 4 {
			continue
		}

		corners[tID] = []*BorderInfo{}
	}

	for _, info := range borders {
		if len(info) != 1 {
			continue
		}

		tID := info[0].TileID
		if _, ok := corners[tID]; ok {
			corners[tID] = append(corners[tID], info[0])
		}
	}

	for cID := range corners {
		for _, fy := range []bool{true, false} {
			for _, fx := range []bool{true, false} {
				builder := NewBuildState(cID, fx, fy, corners, tiles, borders)

				madeMap, ok := builder.build()
				if ok {
					return madeMap
				}
			}
		}
	}

	return ""
}
