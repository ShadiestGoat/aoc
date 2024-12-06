package day12

import (
	"math"
	"strconv"
	"strings"
)

type State struct {
	XY

	Direction int
}

type XY struct {
	X, Y int
}

func (s *XY) getXY(ins rune) *int {
	switch ins {
	case 'S', 'N':
		return &s.Y
	case 'W', 'E':
		return &s.X
	}

	return nil
}

func (s *State) getXY(ins rune) *int {
	if v := s.XY.getXY(ins); v != nil {
		return v
	}

	return &s.Direction
}

func (s *State) FixDir() {
	for s.Direction < 0 {
		s.Direction += 360
	}
	for s.Direction >= 360 {
		s.Direction %= 360
	}
}

func parseLine(inp string) (rune, int) {
	n, _ := strconv.Atoi(inp[1:])

	return rune(inp[0]), n
}

func parseDir(ins rune, v int) int {
	switch ins {
	case 'S', 'W', 'L':
		return -v
	}

	return v
}

func getInsFromDirection(dir int) rune {
	switch dir {
	case 0:
		return 'N'
	case 90:
		return 'E'
	case 180:
		return 'S'
	case 270:
		return 'W'
	}

	return 0
}

func Solve1(inp string) any {
	s := &State{
		Direction: 90,
	}

	for _, v := range strings.Split(inp, "\n") {
		ins, n := parseLine(v)

		var ptr *int

		if ins != 'F' {
			n = parseDir(ins, n)
			ptr = s.getXY(ins)
		} else {
			dir := getInsFromDirection(s.Direction)

			n = parseDir(dir, n)
			ptr = s.getXY(dir)
		}

		*ptr += n

		s.FixDir()
	}

	return int(math.Abs(float64(s.Y)) + math.Abs(float64(s.X)))
}

func Solve2(inp string) any {
	boat := XY{}
	waypoint := XY{
		X: 10,
		Y: 1,
	}

	for _, v := range strings.Split(inp, "\n") {
		ins, n := parseLine(v)
		n = parseDir(ins, n)

		switch ins {
		case 'F':
			boat.X += waypoint.X * n
			boat.Y += waypoint.Y * n
		case 'L', 'R':
			switch n {
			case 90, -270:
				waypoint.X, waypoint.Y = waypoint.Y, -waypoint.X
			case 180, -180:
				waypoint.X *= -1
				waypoint.Y *= -1
			case 270, -90:
				waypoint.X, waypoint.Y = -waypoint.Y, waypoint.X
			}
		default:
			ptr := waypoint.getXY(ins)
			*ptr += n
		}
	}

	return int(math.Abs(float64(boat.Y)) + math.Abs(float64(boat.X)))
}
