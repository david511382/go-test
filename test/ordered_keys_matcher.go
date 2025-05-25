package test

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"
)

type orderedKeysMatcher[T any, Key constraints.Ordered] struct {
	*CompMatcher[[]T]
	orderKeyGetter func(T) Key
}

func NewOrderedKeysMatcher[T any, Key constraints.Ordered](orderKeyGetter func(T) Key, want ...T) *orderedKeysMatcher[T, Key] {
	res := &orderedKeysMatcher[T, Key]{
		CompMatcher:    NewCompMatcher(want),
		orderKeyGetter: orderKeyGetter,
	}
	res.ErrMsgMatcherBase = NewErrMsgMatcherBase(res)
	return res
}

func (m *orderedKeysMatcher[T, Key]) matches(x any) (ok bool, msg string) {
	got, ok := x.([]T)
	if !ok {
		var w T
		msg = fmt.Sprintf("parse %s to %s fail", reflect.TypeOf(x).String(), reflect.TypeOf(w).String())
		return false, msg
	}

	sort.Slice(got, func(i, j int) bool {
		iv := m.orderKeyGetter(got[i])
		jv := m.orderKeyGetter(got[j])
		return iv < jv
	})

	return m.CompMatcher.matches(got)
}
