package test

import (
	"fmt"
	"strconv"
)

type TestCase[Args any, Migrations any, Wants any] struct {
	Name       string
	Args       Args
	Migrations Migrations
	Wants      Wants
}

func NewTestCase[Args any, Migrations any, Wants any](args Args, migrations Migrations, wants Wants) *TestCase[Args, Migrations, Wants] {
	return NewTestCaseWithName("0", args, migrations, wants)
}

func NewTestCaseWithName[Args any, Migrations any, Wants any](name string, args Args, migrations Migrations, wants Wants) *TestCase[Args, Migrations, Wants] {
	return &TestCase[Args, Migrations, Wants]{
		name,
		args,
		migrations,
		wants,
	}
}

func (tc *TestCase[Args, Migrations, Wants]) MakeTestCases() []TestCase[Args, Migrations, Wants] {
	return []TestCase[Args, Migrations, Wants]{
		{
			tc.Name,
			tc.Args,
			tc.Migrations,
			tc.Wants,
		},
	}
}

type TestCases[Args any, Migrations any, Wants any] []ITestSuit[Args, Migrations, Wants]

func (tss TestCases[Args, Migrations, Wants]) MakeTestCases() []TestCase[Args, Migrations, Wants] {
	tcs := make([]TestCase[Args, Migrations, Wants], 0, len(tss))
	for _, ts := range tss {
		tcs = append(tcs, ts.MakeTestCases()...)
	}

	for i, tc := range tcs {
		tcName := tc.Name
		if tcName == "" {
			tcName = strconv.Itoa(i)
		}
		tcs[i].Name = tcName
	}
	return tcs
}

type TestCasesWithName[Args any, Migrations any, Wants any] struct {
	name  string
	cases []ITestSuit[Args, Migrations, Wants]
}

func NewTestCasesWithName[Args any, Migrations any, Wants any](name string, cases ...ITestSuit[Args, Migrations, Wants]) *TestCasesWithName[Args, Migrations, Wants] {
	return &TestCasesWithName[Args, Migrations, Wants]{
		name,
		cases,
	}
}

func (tcwn TestCasesWithName[Args, Migrations, Wants]) MakeTestCases() []TestCase[Args, Migrations, Wants] {
	tss := tcwn.cases
	tcs := make([]TestCase[Args, Migrations, Wants], 0, len(tss))
	for _, ts := range tss {
		tcs = append(tcs, ts.MakeTestCases()...)
	}

	tcNamePrifix := tcwn.name
	for i, tc := range tcs {
		tcName := tc.Name
		if tcName == "" {
			tcName = strconv.Itoa(i)
		}
		tcs[i].Name = fmt.Sprintf("%s-%s", tcNamePrifix, tcName)
	}
	return tcs
}
