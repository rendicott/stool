package harness

import "github.com/gin-gonic/gin"

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
func (i *FakeRunner) RunAllTests(path string) (string, error) {
	res := "my path is " + path
	return res, nil
}
