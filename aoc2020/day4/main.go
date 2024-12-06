package day4

import (
	"strconv"
	"strings"
)

func validateInt(v, min, max int) bool {
	return v >= min && v <= max
}

func parseValidateInt(v string, min, max int) bool {
	n, err := strconv.Atoi(v)
	if err != nil {
		return false
	}

	return validateInt(n, min, max)
}

func validationWrapper(min, max int) func(string) bool {
	return func(s string) bool {
		return parseValidateInt(s, min, max)
	}
}

var FIELD_VALIDATION = map[string](func(string) bool){
	"byr": validationWrapper(1920, 2002),
	"iyr": validationWrapper(2010, 2020),
	"eyr": validationWrapper(2020, 2030),
	"hgt": func(v string) bool {
		if len(v) <= 2 {
			return false
		}

		unit := v[len(v)-2:]
		if unit != "cm" && unit != "in" {
			return false
		}

		n, err := strconv.Atoi(v[:len(v)-2])
		if err != nil {
			return false
		}

		if unit == "cm" {
			return validateInt(n, 150, 193)
		}

		return validateInt(n, 59, 76)
	},
	"hcl": func(v string) bool {
		if len(v) != 7 {
			return false
		}
		if v[0] != '#' {
			return false
		}

		for _, r := range v[1:] {
			valid := (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')

			if !valid {
				return false
			}
		}

		return true
	},
	"ecl": func(v string) bool {
		return map[string]bool{
			"amb": true,
			"blu": true,
			"brn": true,
			"gry": true,
			"grn": true,
			"hzl": true,
			"oth": true,
		}[v]
	},
	"pid": func(v string) bool {
		if len(v) != 9 {
			return false
		}

		for _, r := range v {
			if r < '0' || r > '9' {
				return false
			}
		}

		return true
	},
}

func parseInput(inp string) []map[string]string {
	d := strings.Split(inp, "\n\n")

	allPassports := make([]map[string]string, len(d))

	for _, cfg := range d {
		fields := strings.Fields(cfg)
		p := map[string]string{}

		for _, f := range fields {
			spl := strings.SplitN(f, ":", 2)
			p[spl[0]] = spl[1]
		}

		allPassports = append(allPassports, p)
	}

	return allPassports
}

func Solve1(inp string) any {
	d := parseInput(inp)
	amt := len(d)

	for _, passport := range d {
		for c := range FIELD_VALIDATION {
			if _, ok := passport[c]; !ok {
				amt--
				break
			}
		}
	}

	return amt
}

func Solve2(inp string) any {
	d := parseInput(inp)
	amt := len(d)

	for _, passport := range d {
		for c, valid := range FIELD_VALIDATION {
			if v, ok := passport[c]; !ok || !valid(v) {
				amt--
				break
			}
		}
	}

	return amt
}
