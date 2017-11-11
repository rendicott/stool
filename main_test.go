package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Context("When we run the program", func() {
		XIt("Checks to see if inspec is installed", func() {
			err := verifyEnvironment()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
