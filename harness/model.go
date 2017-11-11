package harness

type TestRunner interface {
	RunAllTests() (string, error)
}

type InspecRunner struct {
}

// todo have this return formatted data
func (i *InspecRunner) RunAllTests() (string, error) {
	return "", nil
}
