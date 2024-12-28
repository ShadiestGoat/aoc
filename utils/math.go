package utils

func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	return a | b
}

func LCM(nums ...int) int {
	lcm := 1

	for _, v := range nums {
		lcm = (v * lcm) / GCD(lcm, v)
	}

	return lcm
}