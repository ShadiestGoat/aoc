package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shadiestgoat/aoc/solvers"
	"github.com/shadiestgoat/aoc/utils"
)

//go:generate go run -C ./template . ..

var (
	FLAG_OFFLINE = flag.Bool("o", false, "Offline (add a cur_day file) <3")
)

func init() {
	flag.Parse()
}

func main() {
	if solvers.YEAR == "" {
		panic("Grr and brr bad build... see scripts/run.sh")
	}

	cookie := os.Getenv("COOKIE")
	if !*FLAG_OFFLINE {
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
	}

	args := flag.Args()
	if len(args) == 0 || len(args) > 2 {
		panic("Grr and brr bad arg amt")
	}

	d, err := strconv.Atoi(args[0])
	if err != nil {
		panic("Grr and brr -- day isn't right, silly :3")
	}

	p := 1
	if len(args) == 2 {
		part, err := strconv.Atoi(args[1])
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
	var inp string

	if *FLAG_OFFLINE {
		tInp, err := os.ReadFile("cur_day")
		utils.PanicIfErr(err, "reading cur_day file")
		inp = string(tInp)
	} else {
		inp = fetchInput(d, cookie)
	}

	fmt.Println("---> Fetched")

	fmt.Println("(2/2) Solving...")

	t := time.Now()
	ans := solve(strings.TrimSpace(inp))
	dur := time.Since(t)

	fmt.Printf("---> Solved (%v)\n", dur)
	fmt.Println("<===================================>")

	fmt.Println(ans)
}
