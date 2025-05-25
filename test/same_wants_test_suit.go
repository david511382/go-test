package test

import (
	"fmt"
	"strconv"
)

type SameWantsTestSuit[Args any, Migrations any, Wants any] struct {
	Name           string
	ArgsMigrations []ArgsMigrations[Args, Migrations]
	Wants          Wants
}

type ArgsMigrations[Args any, Migrations any] struct {
	Name       string
	Args       Args
	Migrations Migrations
}

func NewSameWantsTestSuit[Args any, Migrations any, Wants any](argsMigrations []ArgsMigrations[Args, Migrations], wants Wants) *SameWantsTestSuit[Args, Migrations, Wants] {
	return NewSameWantsTestSuitWithName(
		"",
		argsMigrations,
		wants,
	)
}

func NewSameWantsTestSuitWithName[Args any, Migrations any, Wants any](name string, argsMigrations []ArgsMigrations[Args, Migrations], wants Wants) *SameWantsTestSuit[Args, Migrations, Wants] {
	return &SameWantsTestSuit[Args, Migrations, Wants]{
		name,
		argsMigrations,
		wants,
	}
}

func (ts *SameWantsTestSuit[Args, Migrations, Wants]) MakeTestCases() []TestCase[Args, Migrations, Wants] {
	result := make([]TestCase[Args, Migrations, Wants], len(ts.ArgsMigrations))
	for i, v := range ts.ArgsMigrations {
		tcName := v.Name
		if tcName == "" {
			tcName = strconv.Itoa(i)
		}

		if ts.Name != "" {
			tcName = fmt.Sprintf("%s-%s", ts.Name, tcName)
		}

		result[i] = *NewTestCaseWithName(tcName, v.Args, v.Migrations, ts.Wants)
	}
	return result
}
