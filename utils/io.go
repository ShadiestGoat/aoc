package utils

import (
	"encoding/json"
	"fmt"
)

type JSONAnswer struct {
	V any
}

func (a JSONAnswer) String() string {
	v, err := json.MarshalIndent(a.V, "", "\t")
	PanicIfErr(err, "encoding answer to str")

	return string(v)
}

// Prints v on new lines as json
func PrintJSON(args ...any) {
	for _, a := range args {
		fmt.Println(JSONAnswer{V: a})
	}
}

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type IteratorFunc[T any] func() (T, bool)

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
