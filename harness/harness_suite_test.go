package harness_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHarness(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Harness Suite")
}
