package utils

import (
	"math"
	"strconv"
)

type XY [2]int

var (
	DIR_UP    = XY{0, -1}
	DIR_DOWN  = XY{0, 1}
	DIR_LEFT  = XY{-1, 0}
	DIR_RIGHT = XY{1, 0}

	DIR_DIAG_NE = XY{1, -1}
	DIR_DIAG_SE = XY{1, 1}
	DIR_DIAG_SW = XY{-1, 1}
	DIR_DIAG_NW = XY{-1, -1}

	ALL_DIRECT_DIRS = []XY{
		DIR_UP,
		DIR_RIGHT,
		DIR_DOWN,
		DIR_LEFT,
	}
	ALL_DIAGONALS = []XY{
		DIR_DIAG_NE,
		DIR_DIAG_SE,
		DIR_DIAG_SW,
		DIR_DIAG_NW,
	}
	ALL_DIRS = []XY{
		DIR_UP,
		DIR_DIAG_NE,
		DIR_RIGHT,
		DIR_DIAG_SE,
		DIR_DOWN,
		DIR_DIAG_SW,
		DIR_LEFT,
		DIR_DIAG_NW,
	}
)

func getSize[T []E | string, E any](m []T) XY {
	return XY{len(m[0]), len(m)}
}

func GetSizeString(m []string) XY {
	return getSize[string, any](m)
}

func GetSize[T any](v [][]T) XY {
	return getSize[[]T, T](v)
}

func (c XY) Add(c2 XY) XY {
	return [2]int{c[0] + c2[0], c[1] + c2[1]}
}

func (c XY) MulCoord(c2 XY) XY {
	return [2]int{c[0] * c2[0], c[1] * c2[1]}
}

func (c XY) Mul(v int) XY {
	return [2]int{c[0] * v, c[1] * v}
}

func (c XY) Copy() XY {
	return XY{c[0], c[1]}
}

func unit(v int) int {
	if v == 0 {
		return 0
	} else if v < 0 {
		return -1
	}

	return 1
}

func (c XY) Unit() XY {
	return XY{unit(c[0]), unit(c[1])}
}

// Checks if c is outside of a box size {size}. This assumes the box is 0, 0 -> point {size}.
func (c XY) OutOfBounds(size XY) bool {
	return c[0] < 0 || c[1] < 0 || c[0] >= size[0] || c[1] >= size[1]
}

func (c XY) IsAtOrigin() bool {
	return c[0] == 0 && c[1] == 0
}

func (c XY) Abs() XY {
	return XY{
		int(math.Abs(float64(c[0]))),
		int(math.Abs(float64(c[1]))),
	}
}

func (c XY) ManhattanDistance() int {
	abs := c.Abs()

	return abs[0] + abs[1]
}

func (c XY) ManhattanDistanceTo(c2 XY) int {
	return int(math.Abs(float64(c[0] - c2[0])) + math.Abs(float64(c[1] - c2[1])))
}

// Used to rate in case of a direction
// Only works for the case of (-1, 1) directions.
// Use a multiplication vector to do the rest
// The n is the the amount of times to rotate clockwise (-1 means counterclockwise)
func (c XY) RotateUnitVector(n int) XY {
	deg := float64(n) * math.Pi/4

	cos := math.Cos(deg)
	sin := math.Sin(deg)

	return XY{
		int(math.Round(float64(c[0])*cos - float64(c[1])*sin)),
		int(math.Round(float64(c[0])*sin + float64(c[1])*cos)),
	}
}

// Returns quadrant this coord is in. Uses -y as up.
// 3 0
// 2 1
func (c XY) Quadrant() int {
	switch c.Unit() {
	case XY{0, -1}, XY{1, -1}:
		return 0
	case XY{1, 0}, XY{1, 1}:
		return 1
	case XY{0, 1}, XY{-1, 1}:
		return 2
	case XY{-1, 0}, XY{-1, -1}:
		return 3
	}

	return -1
}

func (c XY) String() string {
	return strconv.Itoa(c[0]) + "," + strconv.Itoa(c[1])
}

func XYFromArr(arr []int) XY {
	return XY{arr[0], arr[1]}
}
