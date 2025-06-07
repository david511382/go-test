package test

import (
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"
)

type errBuf struct {
	msg string
}

func (m *errBuf) Errorf(format string, args ...any) {
	m.msg = fmt.Sprintf(format, args...)
}

type CompMatcher[Want any] struct {
	*ErrMsgMatcherBase
	want Want
}

// Compare gomock function arg by testify
// Ex:
//
//	var t *testing.T
//	mockCtl := gomock.NewController(t)
//	mockObj := NewMockIGoMockObj(mockCtl)
//
//	mockObj.EXPECT().Run(
//	    test.NewCompMatcher(
//	        &Want{},
//	    ),
//	)
func NewCompMatcher[Want any](want Want) *CompMatcher[Want] {
	res := &CompMatcher[Want]{
		want: want,
	}
	res.ErrMsgMatcherBase = NewErrMsgMatcherBase(res)
	return res
}

func (m *CompMatcher[Want]) matches(x any) (ok bool, msg string) {
	got, ok := x.(Want)
	if !ok {
		var w Want
		msg = fmt.Sprintf("parse %s to %s fail", reflect.TypeOf(x).String(), reflect.TypeOf(w).String())
		return false, msg
	}

	buf := &errBuf{}
	ok = assert.Equal(buf, m.want, got)
	msg = buf.msg
	return
}

func (m *CompMatcher[Want]) String() string {
	return reflect.TypeOf(new(Want)).String()
}
