package harness

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAllTests(c *gin.Context) {
	provider := c.MustGet("runner").(TestRunner)
	result, err := provider.RunAllTests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	}
}
