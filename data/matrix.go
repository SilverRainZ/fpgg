package data

type Matrix[T any] struct {
	m [][]T
}

func MatrixFromValue[T any](m [][]T) *Matrix[T] {
	return &Matrix[T]{m: m}
}

func (m *Matrix[T]) Seek(i Pair[int, int]) Maybe[T] {
	if i.First >= len(m.m) {
		return Nothing[T]()
	}
	if i.Second >= len(m.m[i.First]) {
		return Nothing[T]()
	}
	return Just(m.m[i.First][i.Second])
}

func (m *Matrix[T]) Iter() Iter[Iter[T]] {
	s := sliceIter[[]T]{
		s:     m.m,
		delta: 1,
	}
	return &matrixIter[T]{s}
}

func (m *Matrix[T]) RevIter() Iter[Iter[T]] {
	s := sliceIter[[]T]{
		s:     m.m,
		i:     len(m.m) - 1,
		delta: -1,
	}
	return &matrixIter[T]{s}
}

// TODO
func (m *Matrix[T]) Len() int {
	return len(m.m)
}

var (
	_ Iter[Iter[int]] = &matrixIter[int]{}
)

type matrixIter[T any] struct {
	sliceIter[[]T]
}

func (i *matrixIter[T]) Next() Maybe[Iter[T]] {
	maybeS := i.sliceIter.Next()
	if !maybeS.Ok() {
		return Nothing[Iter[T]]()
	}
	var next Iter[T]
	next = &sliceIter[T]{
		s:     maybeS.Just(),
		delta: i.delta,
	}
	return Just(next)
}
