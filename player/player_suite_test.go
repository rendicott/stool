package player_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlayer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Player Suite")
}
