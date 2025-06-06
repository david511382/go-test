package test

type IMatcher interface {
	matches(x any) (ok bool, msg string)
	// String describes what the matcher matches.
	String() string
}
