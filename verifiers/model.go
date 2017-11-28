package verifiers

type Verifier interface {
	Setup(string) error
	Check(string) (string, error)
}

type TestCase struct {
	Name   string
	Result bool
}

type TestSuite struct {
	Name  string
	Tests []TestCase
}
