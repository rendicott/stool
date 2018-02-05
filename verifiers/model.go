package verifiers

type Verifier interface {
	Setup(string) error
	Check(string) (TestSuite, error)
}

type TestCase struct {
	Name    string
	Message string
	Result  bool
}

type TestSuite struct {
	Name     string
	Platform string
	Tests    []TestCase
}
