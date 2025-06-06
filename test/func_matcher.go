package test

import (
	"fmt"
	"reflect"
)

type FuncMatcher[Got any] struct {
	*ErrMsgMatcherBase
	checkFn func(got Got) (ok bool, msg string)
}

func NewFuncMatcher[Got any](checkFn func(got Got) (ok bool, msg string)) *FuncMatcher[Got] {
	res := &FuncMatcher[Got]{
		checkFn: checkFn,
	}
	res.ErrMsgMatcherBase = NewErrMsgMatcherBase(res)
	return res
}

func (m *FuncMatcher[Got]) matches(x any) (ok bool, msg string) {
	got, ok := x.(Got)
	if !ok {
		var w Got
		msg = fmt.Sprintf("parse %s to %s fail", reflect.TypeOf(x).String(), reflect.TypeOf(w).String())
		return false, msg
	}

	return m.checkFn(got)
}

func (m *FuncMatcher[Got]) String() string {
	return reflect.TypeOf(new(Got)).String()
}
