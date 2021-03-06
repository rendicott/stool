package harness

import (
	"github.com/chrisevett/stool/verifiers"
	"github.com/gin-gonic/gin"
)

type FakeRunner struct {
}

func FakeMiddlewareGood() gin.HandlerFunc {
	return func(c *gin.Context) {
		runner := &FakeRunner{}
		c.Set("runner", runner)
		c.Set("path", "fakepath")
		c.Next()
	}
}

func FakeMiddlewareRunnerNotDefined() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("path", "fakepath")
		c.Next()
	}
}

func FakeMiddlewarePathNotDefined() gin.HandlerFunc {
	return func(c *gin.Context) {
		runner := &FakeRunner{}
		c.Set("runner", runner)
		c.Next()
	}
}
func (i *FakeRunner) RunAllTests(path string) (verifiers.TestSuite, error) {
	t := verifiers.TestSuite{}
	t.Name = "fakesuite"
	t.Platform = "fakeplatform"
	t.Tests = []verifiers.TestCase{{Name: "test1", Message: "given when then", Result: true}}

	return t, nil
}
