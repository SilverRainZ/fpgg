package data

import (
	"fmt"
)

type Maybe[T any] struct {
	just T
	ok   bool
}

func Just[T any](just T) Maybe[T] {
	return Maybe[T]{just, true}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func (m Maybe[T]) Must() T {
	if !m.ok {
		panic(fmt.Errorf("nothing in Maybe[%T]", m.just))
	}
	return m.just
}

func (m Maybe[T]) Just() T {
	return m.just
}

func (m Maybe[T]) Ok() bool {
	return m.ok
}

func (m Maybe[T]) Both() (T, bool) {
	return m.just, m.ok
}
