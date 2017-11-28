package harness

import (
	"fmt"

	"github.com/chrisevett/stool/verifiers/inspec"
)

type TestRunner interface {
	RunAllTests(string) (string, error)
}

type InspecRunner struct {
}

// todo have this return formatted data
func (i *InspecRunner) RunAllTests(path string) (string, error) {
	v := inspec.InspecVerifier{}
	err := v.Setup(path)

	if err != nil {
		fmt.Println("error during inspec verifier setup")
		return "", err
	}

	result, err := v.Check(path)

	if err != nil {
		fmt.Println("error during inspec test execution")
		return result, err
	}
	return result, nil

}
