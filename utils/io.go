package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"shadygoat.eu/aoc/solvers"
)

type JSONAnswer struct {
	V any
}

func (a JSONAnswer) String() string {
	v, err := json.MarshalIndent(a.V, "", "\t")
	PanicIfErr(err, "encoding answer to str")

	return string(v)
}

func FetchInput(day int) string {
	req, _ := http.NewRequest(`GET`, fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", solvers.YEAR, day), nil)
	req.Header = map[string][]string{
		"Cookie": {"session=" + os.Getenv("COOKIE")},
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

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type IteratorFunc[T any] func () (T, bool)

type ScannerUtil[T any] struct {
	h IteratorFunc[T]
	v T
}

func (s *ScannerUtil[T]) Next() bool {
	v, ok := s.h()

	s.v = v

	return ok
}

func (s *ScannerUtil[T]) Value() T {
	return s.v
}

func NewScannerUtil[T any](h IteratorFunc[T]) Iterator[T] {
	return &ScannerUtil[T]{
		h: h,
	}
}