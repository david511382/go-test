package test

import (
	"fmt"
)

type ErrMsgMatcherBase struct {
	matcher IMatcher
	msg     string
}

func NewErrMsgMatcherBase(matcher IMatcher) *ErrMsgMatcherBase {
	return &ErrMsgMatcherBase{
		matcher,
		"",
	}
}

func (m *ErrMsgMatcherBase) Matches(x any) bool {
	ok, msg := m.matcher.matches(x)
	m.msg = msg
	return ok
}

// Got is invoked with the received value. The result is used when
// printing the failure message.
func (m *ErrMsgMatcherBase) Got(got any) string {
	if m.msg == "" {
		return fmt.Sprintf("%v", got)
	}
	return m.msg
}
