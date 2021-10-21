package data

type Seekable[K comparable, V any] interface {
	Seek(K) Maybe[V]
	Len() int
}

type Countable interface {
	Len() int
}

type Iterable[T any] interface {
	Iter() Iter[T]
	RevIter() Iter[T]
}

type Iter[T any] interface {
	Next() Maybe[T]
}
