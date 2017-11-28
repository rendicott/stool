package inspec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInspec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inspec Suite")
}
