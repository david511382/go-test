package test

type ITestSuit[Args any, Migrations any, Wants any] interface {
	MakeTestCases() []TestCase[Args, Migrations, Wants]
}
