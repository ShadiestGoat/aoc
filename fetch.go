package main

import (
	"fmt"
	"io"
	"net/http"

	"shadygoat.eu/aoc/solvers"
)

func fetchInput(day int, cookie string) string {
	req, _ := http.NewRequest(`GET`, fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", solvers.YEAR, day), nil)
	req.Header = map[string][]string{
		"Cookie": {"session=" + cookie},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("Failed to fetch input (err): " + err.Error())
	}

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Failed to fetch input (status %v): %v", resp.StatusCode, string(body)))
	}

	return string(body)
}
