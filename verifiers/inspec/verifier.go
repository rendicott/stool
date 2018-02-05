package inspec

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/chrisevett/stool/verifiers"
)

// todo implement setup and check
type InspecVerifier struct {
}

func (i *InspecVerifier) Setup(path string) error {
	err := VerifyInspecInstall()
	if err != nil {
		fmt.Println("Error verifying inspec install in Setup")
		return err
	}
	err = VerifyInspecProfilePath(path)
	if err != nil {
		fmt.Println("Error verifying inspec profile in Setup")
	}
	return err
}

func (i *InspecVerifier) Check(path string) (verifiers.TestSuite, error) {
	///func (i *InspecVerifier) Check(path string) (verifiers.TestSuite, error) {
	r, inspecerr := ExecInspecTests(path)
	// think we need to eat the error here, inspec returns 1 if a failure occurs and we want to ignore it
	// if err != nil {
	// 	fmt.Println("error running tests")
	// 	return verifiers.TestSuite{}, nil
	// }

	err, suite := OutputToSuite(r)
	if err != nil {
		fmt.Println("error converting output")
		return verifiers.TestSuite{}, err
	}
	// need to call the tranform bit here
	// may not need a seperate file tho
	return suite, inspecerr
}

var execCommand = exec.Command

func VerifyInspecProfilePath(path string) error {
	cmd := execCommand("inspec", "check", path)
	err := cmd.Run()
	return err
}

// look at this for formatting https://play.golang.org/p/QUyL3cyTAC
// want to output a reader here so we can create a json payload more easily here
func ExecInspecTests(path string) (io.Reader, error) {
	cmd := execCommand("inspec", "exec", path, "--format=json")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	fmt.Println("stdout", outb.String(), "stderr", errb.String())
	return &outb, err

}

func VerifyInspecInstall() error {
	cmd := execCommand("inspec", "version")
	err := cmd.Run()
	return err
}
