package inspec_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/chrisevett/stool/verifiers"
	. "github.com/chrisevett/stool/verifiers/inspec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func goodTestSuiteFromOut() TestSuite {
	goodtestcases := TestSuite{Name: "sample",
		Platform: "mac_os_x",
		Tests: []TestCase{{Name: "Service sshd should be installed",
			Result: false},
			{Name: "Service sshd should be enabled",
				Result: false},
			{Name: "Service sshd should be running",
				Result: false},
			{Name: "File /tmp should be directory",
				Result: false},
		},
	}

	return goodtestcases
}

var _ = Describe("Given LoadOutput is called", func() {
	Context("when we pass a reader with valid inspec output", func() {
		It("returns a model struct", func() {
			// load test json
			f, _ := os.Open("../out.json")
			err, actual := LoadOutput(f)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual.Version).To(Equal("1.43.8"))
			Expect(actual.Profiles[0].Name).To(Equal("sample"))

		})
	})
	Context("when we pass a reader with invalid input", func() {
		It("returns an error ", func() {
			r := bytes.NewReader([]byte("buttles"))
			err, _ := LoadOutput(r)
			Expect(err).To(HaveOccurred())
		})
	})
})

var _ = Describe("Given InspecOutputToTestSuite is called", func() {
	Context("when we pass a populated model struct", func() {
		It("returns a testcase object", func() {
			// I know this duplicates LoadOutput.
			f, _ := os.Open("../out.json")
			bytes, err := ioutil.ReadAll(f)
			var i InspecOutputModel
			err = json.Unmarshal(bytes, &i)
			err, actual := InspecOutputToTestSuite(&i)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual.Platform).To(Equal("mac_os_x"))
			Expect(actual.Name).To(Equal("sample"))
			Expect(actual.Tests).To(HaveLen(4))
			Expect(actual.Tests[0].Result).To(Equal(false))
		})
	})
})
