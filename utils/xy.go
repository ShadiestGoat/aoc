package utils

import (
	"math"
)

type XY [2]int

var (
	DIR_UP    = XY{0, -1}
	DIR_DOWN  = XY{0, 1}
	DIR_LEFT  = XY{-1, 0}
	DIR_RIGHT = XY{1, 0}

	ALL_DIRECT_DIRS = []XY{
		DIR_UP,
		DIR_RIGHT,
		DIR_DOWN,
		DIR_LEFT,
	}
	ALL_DIRS = []XY{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
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
	deg := float64(n * 45)

	cos := math.Cos(deg)
	sin := math.Sin(deg)

	return XY{
		int(math.Round(float64(c[0])*cos - float64(c[1])*sin)),
		int(math.Round(float64(c[0])*sin + float64(c[1])*cos)),
	}
}

func XYFromArr(arr []int) XY {
	return XY{arr[0], arr[1]}
}