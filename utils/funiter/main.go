package funiter

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
