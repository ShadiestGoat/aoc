package utils

import (
	"strconv"
	"strings"
)

func SplitAndParseInt(inp string, delim string) []int {
	raw := strings.Split(inp, delim)
	p := make([]int, len(raw))

	for i, v := range raw {
		p[i] = ParseInt(v)
	}

	return p
}

func ParseInt(v string) int {
	p, err := strconv.Atoi(v)
	PanicIfErr(err, "parsing '%v' as int", v)

	return p
}