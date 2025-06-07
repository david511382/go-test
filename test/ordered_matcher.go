package test

import (
	"golang.org/x/exp/constraints"
)

type orderedMatcher[T constraints.Ordered] struct {
	*orderedKeysMatcher[T, T]
}

func NewOrderedMatcher[T constraints.Ordered](want ...T) *orderedMatcher[T] {
	return &orderedMatcher[T]{
		orderedKeysMatcher: NewOrderedKeysMatcher(
			func(t T) T {
				return t
			},
			want...,
		),
	}
}
