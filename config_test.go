package main

import (
	"errors"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ValidReader struct {
}

func (v ValidReader) Read(path string) ([]byte, error) {
	filetobytes := `verifier: inspec
testpath: /tmp/profile`

	return []byte(filetobytes), nil
}

type BadFileReader struct {
}

func (v BadFileReader) Read(path string) ([]byte, error) {
	return nil, errors.New("Welp looks like i done goofed")
}

var _ = Describe("Config", func() {
	Context("when I call load config with a well formed config file", func() {
		It("returns a populated config object", func() {
			tmpfile, _ := ioutil.TempFile("", "example")
			defer os.Remove(tmpfile.Name())
			testConfig := Config{Verifier: "inspec", TestPath: "/tmp/profile"}
			rdr := ValidReader{}
			result, _ := LoadConfig(tmpfile.Name(), rdr)

			Expect(result).To(Equal(testConfig))
		})
	})
	Context("when I call load config and the file doesnt exist", func() {
		It("returns an error", func() {
			rdr := BadFileReader{}
			_, err := LoadConfig("/doesnt/matter", rdr)

			Expect(err).Should(HaveOccurred())
		})
	})
	Context("when I call load config and the file can't be read", func() {
		It("returns an error", func() {
			rdr := BadFileReader{}
			_, err := LoadConfig("/doesnt/matter", rdr)

			Expect(err).Should(HaveOccurred())
		})
	})
	Context("when I call FileExists", func() {
		It("does not return an error when the file exists", func() {
			tmpfile, _ := ioutil.TempFile("", "example")
			defer os.Remove(tmpfile.Name())
			err := FileExists(tmpfile.Name())
			Expect(err).ToNot(HaveOccurred())
		})
		It("returns an error when the file doesnt exist", func() {
			err := FileExists("/nothing/nowhere")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when I call parseconfig with a correct config file", func() {

		goodconfig := `verifier: inspec
testpath: /tmp/profile`

		testConfig := Config{Verifier: "inspec", TestPath: "/tmp/profile"}

		It("returns a populated config object", func() {
			result, _ := ParseConfigFile([]byte(goodconfig))
			Expect(result).To(Equal(testConfig))

		})
	})
	Context("when I call parseconfig with an incorrect config file", func() {
		It("returns an error when verifier is missing", func() {
			badconfig := `vrifier: inspec
testpath: /tmp/profile`
			_, err := ParseConfigFile([]byte(badconfig))
			Expect(err).To(HaveOccurred())
		})
		It("returns an error when testpath is missing", func() {
			badconfig := `verifier: inspec
tstpath: /tmp/profile`
			_, err := ParseConfigFile([]byte(badconfig))
			Expect(err).To(HaveOccurred())
		})
	})
})
