package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Suite")
}
