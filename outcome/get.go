package outcome

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OutcomeIndex(c *gin.Context) {
	outcomes, err := GetOutcomes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": outcomes})
	}
}

func ShowOutcome(c *gin.Context) {
	outcomeId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		o := Outcome{Id: outcomeId}
		o, err = o.GetOutcome()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": o.Id,
				"GameId":   o.GameId,
				"PlayerId": o.PlayerId,
				"Result":   o.Result,
				"Score":	o.Score,
				"Date":		o.Date})
		}
	}
}
