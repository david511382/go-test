package test

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type anyOfOrderedMatcher[T constraints.Ordered] map[T]any

func NewAnyOfOrderedMatcher[T constraints.Ordered](ss ...T) anyOfOrderedMatcher[T] {
	res := make(anyOfOrderedMatcher[T])
	for _, s := range ss {
		res[s] = nil
	}
	return res
}

func (m anyOfOrderedMatcher[T]) Matches(x any) bool {
	s, ok := x.(T)
	if !ok {
		return false
	}

	_, exist := m[s]
	return exist
}

func (m anyOfOrderedMatcher[T]) String() string {
	ss := make([]T, 0)
	for s := range m {
		ss = append(ss, s)
	}
	return fmt.Sprintf("is one of %v", ss)
}
