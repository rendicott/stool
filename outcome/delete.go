package outcome

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteOutcome(c *gin.Context) {
	outcomeId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		o := Outcome{Id: outcomeId}
		err = o.DeleteOutcome()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": o.Id,
				"GameId":   o.GameId,
				"PlayerId": o.PlayerId,
				"Win":      o.Win})
		}
	}
}
