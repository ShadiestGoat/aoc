package xy

import (
	"math"
	"strconv"

	"github.com/shadiestgoat/aoc/utils/mutils"
)

type XY [2]int

func (c XY) Add(c2 XY) XY {
	return [2]int{c[0] + c2[0], c[1] + c2[1]}
}

func (c XY) MulCoord(c2 XY) XY {
	return [2]int{c[0] * c2[0], c[1] * c2[1]}
}

func (c XY) Mul(v int) XY {
	return [2]int{c[0] * v, c[1] * v}
}

func (c XY) Unit() XY {
	return XY{mutils.Dir(c[0]), mutils.Dir(c[1])}
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

func MinMaxOfCoords(coords []XY) (XY, XY) {
	min, max := coords[0], coords[0]

	for _, c := range coords[1:] {
		if c[0] < min[0] {
			min[0] = c[0]
		}
		if c[1] < min[1] {
			min[1] = c[1]
		}
		if c[0] > max[0] {
			max[0] = c[0]
		}
		if c[1] > max[1] {
			max[1] = c[1]
		}
	}

	return min, max
}