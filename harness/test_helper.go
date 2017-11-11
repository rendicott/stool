package harness

import "github.com/gin-gonic/gin"

type FakeRunner struct {
}

func FakeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		runner := &FakeRunner{}
		c.Set("runner", runner)
		c.Next()
	}
}

func (i *FakeRunner) RunAllTests() (string, error) {
	return "butts", nil
}
