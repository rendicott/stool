package inspec

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// copying and pasting code from stackoverflow
// #justshitprogrammerthings
// https://npf.io/2015/06/testing-exec-command/
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

// this is going to run in its own process
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	args := os.Args

	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]

	// this one's for twold
	switch cmd {
	case "inspec":
		if args[0] == "check" {
			if args[1] == "../profiles/sample/" {
				fmt.Fprintf(os.Stdout, "GET ME OUT OF THIS WHACK ASS CRYSTAL PRISON")
				os.Exit(0)

			} else {
				fmt.Fprintf(os.Stderr, "invalid profile path ")
				os.Exit(1)
			}
		}
		if args[0] == "exec" {
			fmt.Println("entering exec")
			fmt.Println(args[1])
			fmt.Println(args[2])
			if args[1] == "../profiles/sample/" {
				fmt.Fprintf(os.Stdout, "GET ME OUT OF THIS WHACK ASS CRYSTAL PRISON")
				os.Exit(0)
			} else if args[1] == "../profiles/failsample/" {
				fmt.Fprintf(os.Stdout, "this is sa failing tests")
				os.Exit(1)
			} else {
				fmt.Fprintf(os.Stderr, "invalid profile path ")
				os.Exit(1)
			}
		}
		if args[0] == "version" {
			fmt.Fprintf(os.Stdout, "fake version string from inspec")
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}

var _ = Describe("VerifyInspecProfilePath", func() {
	Context("when I pass a valid path", func() {
		It("returns no error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			err := VerifyInspecProfilePath("../profiles/sample/")
			Expect(err).ShouldNot(HaveOccurred())
		})

	})
	Context("when I pass an invalid path", func() {
		It("returns an error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			err := VerifyInspecProfilePath("/dev/null")
			Expect(err).Should(HaveOccurred())

		})
	})
})

var _ = Describe("ExecInspecTests", func() {
	Context("When we run passing tests ", func() {
		It("returns json results and no error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			_, err := ExecInspecTests("../profiles/sample/")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
	Context("When we run failing tests", func() {
		It("returns json results and an error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			_, err := ExecInspecTests("../profiles/failsample/")
			Expect(err).Should(HaveOccurred())
		})
	})
})

var _ = Describe("ExecInspecTests", func() {
	Context("When we run VerifyInspecInstall", func() {
		It("returns no error if the inspec command executes correctly", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			err := VerifyInspecInstall()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

var _ = XDescribe("really exec inspec ", func() {
	Context("running exec", func() {
		It("returns a string", func() {
			execCommand = exec.Command
			result, _ := ExecInspecTests("../../testprofiles/sample2/")
			Expect(result).ShouldNot(BeEmpty())
		})
	})
})
