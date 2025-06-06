# go-test

Test tool for golang

<p align="left">
	* <a href="#Features">Features</a><br>
    * <a href="#Installation">Installation</a><br>
    * <a href="#Examples">Examples</a><br>
    * <a href="#Dependencies">Dependencies</a><br>
</p>

# Features

* Implement for package gomock Matcher
* Test case architecture

# Installation

## Install with the `go` tool

```bash
go get github.com/david511382/go-test
```

# Examples

## Implement for package gomock Matcher

``` go
var t *testing.T
mockCtl := gomock.NewController(t)
mockObj := NewMockIGoMockObj(mockCtl)

mockObj.EXPECT().Run(
    test.NewCompMatcher(
        &Want{},
    ),
)
```

## Test case architecture

``` go
type args struct {}
type migrations struct {}
type wants struct {}
tests := map[string]test.ITestSuit[args, migrations, wants]{
	// "Test Case"
	"Test Case": test.NewTestCase(
		args{},
		migrations{},
		wants{},
	),
	"Same Result Test Cases": test.NewSameWantsTestSuit(
		[]test.ArgsMigrations[args, migrations]{
			{
				// Default named by index, "Same Result Test Cases-0"
				Args: args{},
				Migrations: migrations{},
			},
			{
				// "Same Result Test Cases-Case1"
				Name: "Case1",
				Args: args{},
				Migrations: migrations{},
			},
		},
		wants{},
	),
	"Same Category Test Cases": test.TestCases[args, migrations, wants]{
		test.TestCases[args, migrations, wants]{
			test.NewTestCaseWithName(
				// "Same Category Test Cases-Case1"
				"Case1",
				args{},
				migrations{},
				wants{},
			),
		},
		test.NewTestCaseWithName(
			// "Same Category Test Cases-Case2"
			"Case2",
			args{},
			migrations{},
			wants{},
		),
	},
}
for testSuitName, testSuit := range tests {
	testSuit := testSuit
	for _, tc := range testSuit.MakeTestCases() {
		tc := tc
		testCaseName := fmt.Sprintf("%s-%s", testSuitName, tc.Name)
		t.Run(testCaseName, func(t *testing.T) {
			// tc.Migrations
			// tc.Args
			// tc.Wants
		})
	}
}
```

# Dependencies

* github.com/golang/mock/gomock v1.16.0
* github.com/stretchr/testify v1.10.0
