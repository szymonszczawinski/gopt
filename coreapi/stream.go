package coreapi

import "iter"

func Map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func Map[I, O any](input iter.Seq[I], transform func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for i := range input {
			if !yield(transform(i)) {
				return
			}
		}
	}
}

func Filter[I any](input iter.Seq[I], filter func(I) bool) iter.Seq[I] {
	return func(yield func(I) bool) {
		for i := range input {
			if filter(i) {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func Stream[I any](input []I) iter.Seq[I] {
	return func(yield func(I) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}
}

func ToSlice[I any](input iter.Seq[I]) []I {
	result := []I{}
	for i := range input {
		result = append(result, i)
	}
	return result
}
