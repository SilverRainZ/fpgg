package fn

import (
	"constraints"
)

func Zero[T comparable](v T) bool {
	var zero T
	return v == zero
}

func NonZero[T comparable](v T) bool {
	var zero T
	return v != zero
}

func Equal[T constraints.Ordered](a, b T) bool {
	return a == b
}

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

func More[T constraints.Ordered](a, b T) bool {
	return a < b
}
