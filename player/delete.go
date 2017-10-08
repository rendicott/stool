package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePlayer(c *gin.Context) {
	dataContext := c.MustGet("Db").(PlayerDLInterface)
	playerId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		err := dataContext.DeletePlayer(playerId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "successful"})
		}
	}

}
