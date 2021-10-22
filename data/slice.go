package data

var (
	_ Seekable[int, int] = &Slice[int]{}
	_ Iterable[int]      = &Slice[int]{}
	_ Countable          = &Slice[int]{}
)

type Slice[T any] struct {
	s []T
}

func FromSlice[T any](s []T) *Slice[T] {
	return &Slice[T]{s: s}
}

func (s *Slice[T]) Seek(i int) Maybe[T] {
	if i < len(s.s) {
		return Just(s.s[i])
	}
	return Nothing[T]()
}

func (s *Slice[T]) Iter() Iter[T] {
	return &sliceIter[T]{
		s:     s.s,
		delta: 1,
	}
}

func (s *Slice[T]) RevIter() Iter[T] {
	return &sliceIter[T]{
		s:     s.s,
		i:     len(s.s) - 1,
		delta: -1,
	}
}

func (s *Slice[T]) Len() int {
	return len(s.s)
}

func (*Slice[T]) Replace(i Iter[T]) *Slice[T] {
	var s []T
	for {
		v := i.Next()
		if !v.Ok() {
			break
		}
		s = append(s, v.Just())
	}
	return &Slice[T]{s: s}
}

var (
	_ Iter[int] = &sliceIter[int]{}
)

type sliceIter[T any] struct {
	s     []T
	i     int
	delta int
}

func (i *sliceIter[T]) Next() Maybe[T] {
	if (i.delta > 0 && i.i < len(i.s)) || (i.delta < 0 && i.i >= 0) {
		v := i.s[i.i]
		i.i += i.delta
		return Just(v)
	}
	return Nothing[T]()
}
