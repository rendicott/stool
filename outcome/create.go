package outcome

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// one problem here, BindJSON doesn't like being passed an id as a string
func CreateOutcome(c *gin.Context) {
	var o Outcome
	if c.BindJSON(&o) == nil {
		if err := o.CreateOutcome(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": o.Id,
				"GameId":   o.GameId,
				"PlayerId": o.PlayerId,
				"Result":   o.Result,
				"Score":	o.Score,
				"Date":		o.Date})
		}
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
}
