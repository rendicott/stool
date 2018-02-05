package harness

import (
	"fmt"

	"github.com/chrisevett/stool/verifiers"
	"github.com/chrisevett/stool/verifiers/inspec"
)

type TestRunner interface {
	RunAllTests(string) (verifiers.TestSuite, error)
}

type InspecRunner struct {
}

// todo have this return formatted data
func (i *InspecRunner) RunAllTests(path string) (verifiers.TestSuite, error) {
	v := inspec.InspecVerifier{}
	err := v.Setup(path)

	if err != nil {
		fmt.Println("error during inspec verifier setup")
		return verifiers.TestSuite{}, err
	}

	result, err := v.Check(path)

	if err != nil {
		fmt.Println("error during inspec test execution")
		return result, err
	}
	return result, nil

}
