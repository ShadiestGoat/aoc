package day20

import (
	"math"
	"slices"
)

// Rotates string s 90 deg, where n is how many times
// String should be a square
func rotateString(lines []string, n int) []string {
	cos := math.Cos(float64(n) * math.Pi / 2)
	sin := math.Sin(float64(n) * math.Pi / 2)

	center := float64(len(lines))/2 - 0.5

	newStr := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		newStr[i] = make([]rune, len(lines))
	}

	for y, l := range lines {
		for x, r := range l {
			tx, ty := float64(x) - center, float64(y) - center

			nx := tx * cos - ty * sin + center
			ny := tx * sin + ty * cos + center
			newStr[int(math.Round(ny))][int(math.Round(nx))] = r
		}
	}

	o := make([]string, len(newStr))
	for i, v := range newStr {
		o[i] = string(v)
	}

	return o
}

func flipX(lines []string) []string {
	ns := make([]string, len(lines))

	for i, l := range lines {
		nl := make([]byte, len(l))
		copy(nl, l)
		slices.Reverse(nl)

		ns[i] = string(nl)
	}

	return ns
}

func flipY(lines []string) []string {
	ns := make([][]rune, len(lines))
	for i := range ns {
		ns[i] = make([]rune, len(lines[i]))
	}

	for x := range lines[0] {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			ns[i][x], ns[j][x] = rune(lines[j][x]), rune(lines[i][x])
		}
	}

	str := []string{}
	for _, v := range ns {
		str = append(str, string(v))
	}

	return str
}
