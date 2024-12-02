package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: tpl /path/to/aoc/root")
	}

	_, err := os.Stat(os.Args[1] + "/go.mod")
	if err != nil {
		panic("Bad path specified (must be aoc root)")
	}

	dir, err := os.ReadDir(os.Args[1])
	if err != nil {
		panic("Can't read dir: " + err.Error())
	}

	for _, f := range dir {
		if !f.IsDir() {
			continue
		}

		n := f.Name()
		if !strings.HasPrefix(n, "aoc") {
			continue
		}

		if makeYearFile(os.Args[1] + "/" + n) {
			makeWrapperSolver(os.Args[1], n[3:])
		}
	}
}
