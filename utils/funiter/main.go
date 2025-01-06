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

func Scan(inp string, delim string, h ScanFunc) bool {
	li := 0

	for i := range inp {
		if i - li + 1 >= len(delim) && inp[i - len(delim) + 1:i + 1] == delim {
			if h(inp[li:i - len(delim) + 1]) {
				return true
			}

			li = i + 1
		}
	}

	return h(inp[li:])
}

type ScanReduceFunc[T any] func (s string, acc T) T
type ScanFuncTrans[T any] func (v T) bool
type ScanFunc = ScanFuncTrans[string]

func ScanReduce[T any](inp string, delim string, initial T, h ScanReduceFunc[T]) T {
	v := initial

	Scan(inp, delim, func(s string) bool {
		v = h(s, v)

		return false
	})

	return v
}

func NewScanFunc(delim string) func (s string, h ScanFunc) bool {
	return func (s string, h ScanFunc) bool {
		return Scan(s, delim, h)
	}
}

func NewScanTransFunc[T any](delim string, trans func (s string) T) func (s string, h ScanFuncTrans[T]) bool {
	return func (s string, h ScanFuncTrans[T]) bool {
		return Scan(s, delim, func (s string) bool {
			return h(trans(s))
		})
	}
}

func NewScanReduceFunc[T any](delim string, initial T) func (s string, h ScanReduceFunc[T]) T {
	return func (s string, h ScanReduceFunc[T]) T {
		return ScanReduce(s, delim, initial, h)
	}
}

func NewScanSumFunc(delim string) func (s string, h func (s string) int) int {
	return func(s string, h func (s string) int) int {
		return ScanReduce(s, delim, 0, func(s string, acc int) int {
			return acc + h(s)
		})
	}
}
