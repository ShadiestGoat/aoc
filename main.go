package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"shadygoat.eu/aoc/solvers"
	"shadygoat.eu/aoc/utils"
)

//go:generate go run -C ./template . ..

func main() {
	if solvers.YEAR == "" {
		panic("Grr and brr bad build... see scripts/run.sh")
	}

	cookie := os.Getenv("COOKIE")

	if cookie == "" {
		c, err := os.ReadFile(".conf/.cookie")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				panic("Grrrrrr and brrrr... did you forget to put a cookie in?")
			}
			utils.PanicIfErr(err, "reading cookie file")
		}

		cookie = strings.TrimSpace(string(c))
	}

	if len(os.Args) <= 1 || len(os.Args) > 3 {
		panic("Grr and brr bad arg amt")
	}

	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Grr and brr -- day isn't right, silly :3")
	}

	p := 1
	if len(os.Args) == 3 {
		part, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("Grr and brr -- day part isn't right, silly girl :33")
		}
		if part < 1 || part > 2 {
			panic("Grr and brr >:( wt hecc bad part but like really bad")
		}

		p = part
	}

	solve, ok := solvers.Solvers[fmt.Sprintf("%v-%v", d, p)]
	if !ok {
		panic("Grr and brr -- that day isn't created yet, silly woman :3")
	}

	fmt.Println("(1/2) Fetching Input...")
	inp := fetchInput(d, cookie)
	fmt.Println("---> Fetched")

	fmt.Println("(2/2) Solving...")

	t := time.Now()
	ans := solve(strings.TrimSpace(inp))
	dur := time.Since(t)

	fmt.Printf("---> Solved (%v)\n", dur)
	fmt.Println("<===================================>")

	fmt.Println(ans)
}
