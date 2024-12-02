package main

import (
	"fmt"
	"os"
)

const (
	SOLVER_FILE = `//go:build %v

package solvers

import "github.com/shadiestgoat/aoc/aoc%v"

func init() {
	Solvers = aoc%v.Solvers
}
`
)

func makeWrapperSolver(basePath string, year string) {
	os.WriteFile(
		basePath+"/solvers/"+year+".go",
		[]byte(fmt.Sprintf(SOLVER_FILE, year, year, year)),
		0755,
	)
}
