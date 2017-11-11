package verifier

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Verifier interface {
	Check(string) (string, error)
}

type InspecVerifier struct {
}

//func (i *InspecVerifier) Check(profilePath string) (string, error) {

// check profile path for inspec.ymls

// run inspec

// return formatted report
//}
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
