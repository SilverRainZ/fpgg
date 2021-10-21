package fp

type FuncWith1Args[A, R any] func(A) R

func (f FuncWith1Args[A, R]) Parital(a A) R {
	return f(a)
}

type FuncWith2Args[A1, A2, R any] func(A1, A2) R

func Cast[A1, A2, R any](f FuncWith2Args[A1, A2, R]) FuncWith2Args[A1, A2, R] {
	return f
}

func (f FuncWith2Args[A1, A2, R]) Parital(a1 A1) FuncWith1Args[A2, R] {
	return func(a2 A2) R {
		return f(a1, a2)
	}
}

type FuncWith3Args[A1, A2, A3, R any] func(A1, A2, A3) R

func (f FuncWith3Args[A1, A2, A3, R]) Parital(a1 A1) FuncWith2Args[A2, A3, R] {
	return func(a2 A2, a3 A3) R {
		return f(a1, a2, a3)
	}
}
