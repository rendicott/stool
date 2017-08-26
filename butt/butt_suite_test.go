package butt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestButt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Butt Suite")
}
