package day8

// rows, pixels, all layers of pixel
func ParseInput(inp string, mx, my int) ([][][]int, map[int]map[int]int) {
	board := [][][]int{
		{},
	}
	layerToDigit := map[int]map[int]int{
		0: {},
	}

	x, y, l := 0, 0, 0

	for _, r := range inp {
		v := int(r - '0')

		if l == 0 {
			board[y] = append(board[y], []int{v})
		} else {
			board[y][x] = append(board[y][x], v)
		}

		layerToDigit[l][v]++

		x++
		if x >= mx {
			x = 0
			y++
			if y < my {
				board = append(board, [][]int{})
			}
		}
		if y >= my {
			y = 0
			l++
			layerToDigit[l] = map[int]int{}
		}
	}

	return board[:len(board) - 1], layerToDigit
}

func GenericSolve1(inp string, mx, my int) int {
	_, layerData := ParseInput(inp, mx, my)
	
	smallestLayerID := -1
	smallestCount := 0

	for lID, d := range layerData {
		if smallestCount == 0 || d[0] < smallestCount {
			smallestCount = d[0]
			smallestLayerID = lID
		}
	}

	d := layerData[smallestLayerID]

	return d[1] * d[2]
}

func Solve1(inp string) any {
	return GenericSolve1(inp, 25, 6)
}

func Solve2(inp string) any {
	b, _ := ParseInput(inp, 25, 6)

	o := ""

	for _, row := range b {
		if len(row) == 0 {
			// Idk why it prints a bunch of empty rows, but this is a hacky fix so
			break
		}

		o += "\n"

		for _, pixel := range row {
			for _, l := range pixel {
				if l == 2 {
					continue
				}

				o += string(rune(l + '0'))
				break
			}
		}
	}

	return o[1:]
}
