package day2

import (
	"strconv"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type PasswordPolicy struct {
	Min, Max int
	Letter   byte
}

func (p PasswordPolicy) IsValidPassword1(pass string) bool {
	amt := 0

	for _, r := range pass {
		if byte(r) == p.Letter {
			amt++
		}

		if amt > p.Max {
			return false
		}
	}

	return amt >= p.Min
}

func (p PasswordPolicy) IsValidPassword2(pass string) bool {
	if len(pass) < p.Min {
		return false
	}

	min := pass[p.Min-1] == p.Letter

	if len(pass) < p.Max {
		return min
	}

	max := pass[p.Max-1] == p.Letter

	return min != max
}

type DBEntry struct {
	Policy   *PasswordPolicy
	Password string
}

func parseInput(inp string) []*DBEntry {
	pass := []*DBEntry{}

	for _, l := range strings.Split(inp, "\n") {
		if l == "" {
			continue
		}

		spl := strings.Split(l, ": ")
		if len(spl) != 2 {
			panic("Bad split of l '" + l + "'")
		}

		spl2 := strings.Split(spl[0][:len(spl[0])-2], "-")
		min, err := strconv.Atoi(spl2[0])
		utils.PanicIfErr(err, "parsing min in '%v'", spl[0])

		max, err := strconv.Atoi(spl2[1])
		utils.PanicIfErr(err, "parsing max in '%v'", spl[0])

		pass = append(pass, &DBEntry{
			Policy: &PasswordPolicy{
				Min: min,
				Max: max,
				// Good practices only here
				Letter: spl[0][len(spl[0])-1],
			},
			Password: spl[1],
		})
	}

	return pass
}

func Solve1(inp string) any {
	passports := parseInput(inp)
	validAmt := 0

	for _, v := range passports {
		if v.Policy.IsValidPassword1(v.Password) {
			validAmt++
		}
	}

	return validAmt
}

func Solve2(inp string) any {
	passports := parseInput(inp)
	validAmt := 0

	for _, v := range passports {
		if v.Policy.IsValidPassword2(v.Password) {
			validAmt++
		}
	}

	return validAmt
}
