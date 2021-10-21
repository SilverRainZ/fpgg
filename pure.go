package fp

import (
	"github.com/SilverRainZ/fpgg/data"
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
