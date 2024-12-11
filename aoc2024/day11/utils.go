package day11

import "math"


func countDigits(v int) int {
	return int(math.Floor(math.Log10(float64(v)))) + 1
}

func numberLogic(n int) []int {
	if n == 0 {
		return []int{1}
	} else if c := countDigits(n); c % 2 == 0 {
		p1 := math.Pow10(c/2)
		v1 := math.Floor(float64(n)/p1)
		v2 := float64(n) - v1 * p1
	
		return []int{int(v1), int(v2)}
	} else {
		return []int{n * 2024}
	}
}
