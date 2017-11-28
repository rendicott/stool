package inspec

import (
	"bytes"
	"fmt"
	"os/exec"
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

func (i *InspecVerifier) Check(path string) (string, error) {
	result, err := ExecInspecTests(path)
	return result, err
}

var execCommand = exec.Command

func VerifyInspecProfilePath(path string) error {
	cmd := execCommand("inspec", "check", path)
	err := cmd.Run()
	return err
}

// lets forget about checking the pass or fail state at this point

func ExecInspecTests(path string) (string, error) {
	cmd := execCommand("inspec", "exec", path, "--format=json")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	fmt.Println("stdout", outb.String(), "stderr", errb.String())
	return outb.String(), err

}

func VerifyInspecInstall() error {
	cmd := execCommand("inspec", "version")
	err := cmd.Run()
	return err
}
