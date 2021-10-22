package unpure

import (
	"constraints"

	"github.com/SilverRainZ/fpgg/data"
	"github.com/SilverRainZ/fpgg/util"
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

func Fold1[T any](f func(T, T) T, src data.Iter[T]) T {
	return Fold(f, src.Next().Must(), src)
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

// O(1)
func Head[T any](i data.Iter[T]) T {
	return i.Next().Must()
}

// O(N)
func Last[T any](i data.Iter[T]) T {
	prev := i.Next().Must()
	for {
		v := i.Next()
		if !v.Ok() {
			return prev
		}
		prev = v.Just()
	}
}

func Reverse[T any](i data.Iter[T]) data.Iter[T] {
	return data.FromSlice(List(i)).RevIter()
}

func MaxOfOrdered[T constraints.Ordered](i data.Iter[T]) data.Maybe[T] {
	return Max[T](util.Less[T], i)
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

func MinOfOrdered[T constraints.Ordered](i data.Iter[T]) data.Maybe[T] {
	return Min[T](util.Less[T], i)
}

func Min[T any](less func(T, T) bool, i data.Iter[T]) data.Maybe[T] {
	more := func(a, b T) bool { return !less(a, b) }
	return Max[T](more, i)
}

func Take[T any](n int, i data.Iter[T]) data.Iter[T] {
	return Filter(func(v T) bool {
		n--
		return n >= 0
	}, i)
}

func And(i data.Iter[bool]) bool {
	return Fold(func(a, b bool) bool { return a && b }, true, i)
}

func Or(i data.Iter[bool]) bool {
	return Fold(func(a, b bool) bool { return a && b }, true, i)
}

func Any(i data.Iter[bool]) bool {
	return Fold(func(a, b bool) bool { return a && b }, true, i)
}

type concatIter[T any] struct {
	src data.Iter[data.Iter[T]]
	cur data.Iter[T]
}

func (i *concatIter[T]) Next() data.Maybe[T] {
	if i.cur == nil {
		maybeCur := i.src.Next()
		if !maybeCur.Ok() {
			return data.Nothing[T]()
		}
		i.cur = maybeCur.Just()
	}
	v := i.cur.Next()
	if !v.Ok() {
		maybeCur := i.src.Next()
		if !maybeCur.Ok() {
			return data.Nothing[T]()
		}
		i.cur = maybeCur.Just()
		return i.Next()
	}
	return v
}

func Concat[T any](src data.Iter[data.Iter[T]]) data.Iter[T] {
	return &concatIter[T]{src: src}
}
