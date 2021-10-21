package unpure

import (
	"github.com/SilverRainZ/fpgg/data"
)

type mapIter[T1, T2 any] struct {
	f   func(T1) T2
	src data.Iter[T1]
}

func (i *mapIter[T1, T2]) Next() data.Maybe[T2] {
	v := i.src.Next()
	if v.Ok() {
		return data.Just(i.f(v.Just()))
	}
	return data.Nothing[T2]()
}

func Map[T1, T2 any](f func(T1) T2, src data.Iter[T1]) data.Iter[T2] {
	return &mapIter[T1, T2]{f: f, src: src}
}

func Fold[T1, T2 any](f func(T1, T2) T1, init T1, src data.Iter[T2]) T1 {
	res := init
	for {
		v := src.Next()
		if !v.Ok() {
			return res
		}
		res = f(res, v.Just())
	}
}

type filterIter[T any] struct {
	f   func(T) bool
	src data.Iter[T]
}

func (i *filterIter[T]) Next() data.Maybe[T] {
	for {
		v := i.src.Next()
		if !v.Ok() || i.f(v.Just()) {
			return v
		}
	}
}

func Filter[T any](f func(T) bool, src data.Iter[T]) data.Iter[T] {
	return &filterIter[T]{f: f, src: src}
}

func Collect[T any](i data.Iter[T], f func(v T)) {
	for {
		v := i.Next()
		if !v.Ok() {
			return
		}
		f(v.Just())
	}
}

func List[T any](src data.Iter[T]) (dst []T) {
	Collect(src, func(v T) { dst = append(dst, v) })
	return
}

func Head[T any](i data.Iter[T]) data.Maybe[T] {
	return i.Next()
}

// O(N)
func Tail[T any](i data.Iter[T]) data.Maybe[T] {
	prev := data.Nothing[T]()
	for {
		v := i.Next()
		if !v.Ok() {
			return prev
		}
		prev = v
	}
}

func Reverse[T any](i data.Iter[T]) data.Iter[T] {
	return data.FromSlice(List(i)).RevIter()
}

func Max[T any](less func(T, T) bool, i data.Iter[T]) data.Maybe[T] {
	maybeMax := i.Next()
	if !maybeMax.Ok() {
		return maybeMax
	}
	max := maybeMax.Just()
	Collect(i, func(v T) {
		if less(max, v) {
			max = v
		}
	})
	return data.Just(max)
}

func Min[T any](less func(T, T) bool, i data.Iter[T]) data.Maybe[T] {
	more := func(a, b T) bool { return !less(a, b) }
	return Max[T](more, i)
}
