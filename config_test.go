package main

import (
	"errors"
	"flag"
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

// note: i stole this from the test suite for the golang flag package
// if you're reading this their tests are tight af and whenever you're wondering
// about how to test something look at how the golang authors do it
// https://golang.org/src/flag/export_test.go
//
// the issue here was that ginkgo was running ParseConfigPath multiple times
// in a single process which means that it was trying to redefine
// flags which causes a panic.
func ResetFlagsForTesting() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
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

	Context("when I parse the config path from the command line arguments", func() {
		It("returns the contents of the config path if the argument is passed", func() {
			ResetFlagsForTesting()
			// dont want to overwrite a global var for the entire test suite
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = []string{"./stool", "-config=buttles"}
			actual, err := ParseConfigPath()
			Expect(actual).To(Equal("buttles"))
			Expect(err).ToNot(HaveOccurred())
		})
		It("returns the contents of the config path if a different argument is passed", func() {
			// i wrote this because i did something very silly
			ResetFlagsForTesting()
			// dont want to overwrite a global var for the entire test suite
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = []string{"./stool", "-config=trucks"}
			actual, err := ParseConfigPath()
			Expect(actual).To(Equal("trucks"))
			Expect(err).ToNot(HaveOccurred())
		})
		It("returns an error if the config argument is not passed", func() {
			ResetFlagsForTesting()
			// dont want to overwrite a global var for the entire test suite
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// the first arg is the process name so this is testing no arguments passed
			os.Args = []string{"processnamenotanarg"}
			actual, err := ParseConfigPath()
			Expect(actual).To(Equal("ERROR"))
			Expect(err).To(HaveOccurred())
		})
	})

})
