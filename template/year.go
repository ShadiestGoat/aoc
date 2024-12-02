package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	BASE_PKG = "shadygoat.eu/aoc/"
	FILE_TPL = `package %v
// THIS FILE IS AUTO GENERATED, DO NOT EDIT MANUALLY!

import (
%v
)

var Solvers = map[string]func (string) any {
%v
}
`
	EMPTY_SOLVER = `func Solve\d+?\(.+? string\) any \{\n\s+?return nil\n\}`
)

var (
	REG_SOLVER = regexp.MustCompile(`func Solve\d+?\(.+? string\) any`)
	REG_EMPTY_SOLVER = regexp.MustCompile(EMPTY_SOLVER)
)

func makeYearFile(p string) bool {
	dir, err := os.ReadDir(p)
	if err != nil {
		panic("Can't read base dir: " + err.Error())
	}

	days := []string{}
	solverCount := map[string]int{}

	for _, d := range dir {
		if !d.IsDir() {
			continue
		}

		n := d.Name()
		if !strings.HasPrefix(n, "day") {
			continue
		}

		m, err := os.ReadFile(p + "/" + n + "/main.go")
		if err != nil {
			continue
		}

		c := len(REG_SOLVER.FindAll(m, -1)) - len(REG_EMPTY_SOLVER.FindAll(m, -1))
		if c != 0 {
			solverCount[n] = c
			days = append(days, n)
		}
	}

	if len(days) == 0 {
		return false
	}

	y := path.Base(p)
	importList := []string{}
	solveList := []string{}

	for _, d := range days {
		importList = append(importList, "\t\"" + BASE_PKG + y + "/" + d + `"`)
		s := ""

		dayC := d[3:]
		for i := 1; i <= solverCount[d]; i++ {
			s += fmt.Sprintf("\t\"%s-%d\": %s.Solve%d,\n", dayC, i, d, i)
		}

		solveList = append(solveList, s[:len(s) - 1])
	}

	cont := fmt.Sprintf(FILE_TPL, y, strings.Join(importList, "\n"), strings.Join(solveList, "\n\n"))
	os.WriteFile(p + "/solvers.go", []byte(cont), 0755)

	return true
}
