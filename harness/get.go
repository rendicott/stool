package harness

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAllTests(c *gin.Context) {
	runner, runnerExists := c.Get("runner")
	path, pathExists := c.Get("path")
	if runnerExists == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": "runner does not exist"})
	} else if pathExists == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": "path does not exist"})
	} else {
		// need to use type casting here so go can handle the interface properly
		r := runner.(TestRunner)
		p := path.(string)
		result, err := r.RunAllTests(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": result})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
		}
	}
}
