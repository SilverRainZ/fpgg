package fp

import (
	"github.com/SilverRainZ/fpgg/data"
	"github.com/SilverRainZ/fpgg/unpure"
)

func Head[T any](i data.Iterable[T]) data.Maybe[T] {
	return i.Iter().Next()
}

func Tail[T any](i data.Iterable[T]) data.Maybe[T] {
	return i.RevIter().Next()
}

type reverseIterable[T any] struct {
	i data.Iterable[T]
}

func (i *reverseIterable[T]) Iter() data.Iter[T] {
	return i.i.RevIter()
}

func (i *reverseIterable[T]) RevIter() data.Iter[T] {
	return i.i.Iter()
}

func Reverse[T any](i data.Iterable[T]) data.Iterable[T] {
	return &reverseIterable[T]{i}
}

func Len(c data.Countable) int {
	return c.Len()
}

func FMap[T1, T2 any](f func(T1) T2, src data.Functor[T]) data.Functor[T] {
	dstIter := unpure.Map(f, src.Iter())
	return src.Replace(dstIter)
}
