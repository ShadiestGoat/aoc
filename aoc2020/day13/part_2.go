package day13

import (
	"strconv"
	"strings"
)

// Alright yeah so we're gonna do the ol chinese remainder theorem
// thanks geeks4geeks btw <3

// This maybe can be optimized to use a thing specific to primes or maybe not idk im tired and this is a solution so
func inverseModulo(a, m int) int {
	if m == 1 {
		return 0
	}

	m0, x0, x1 := m, 0, 1

	for a > 1 {
		q := a / m

		m, a = a%m, m

		x0, x1 = x1-q*x0, x0
	}

	if x1 < 0 {
		x1 += m0
	}

	return x1
}

func GenericCRA(nums []int, rems []int, prod int) int {
	resp := 0

	for i, n := range nums {
		portion := prod / n
		resp += rems[i] * inverseModulo(portion, n) * portion
	}

	return resp % prod
}

func Product(nums []int) int {
	prod := 1

	for _, v := range nums {
		prod *= v
	}

	return prod
}

func Solve2(inp string) any {
	parsed := strings.Split(inp, "\n")[1]

	nums := []int{}
	rems := []int{} // We need to store remainders due to shifting index (bc we skip x)
	prod := 1

	for i, raw := range strings.Split(parsed, ",") {
		if raw == "x" {
			continue
		}

		v, _ := strconv.Atoi(raw)
		nums = append(nums, v)
		prod *= v

		rem := v - i
		for rem < 0 {
			rem += v
		}

		rems = append(rems, rem)
	}

	return GenericCRA(nums, rems, prod)
}
