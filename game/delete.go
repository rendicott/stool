package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteGame(c *gin.Context) {
	dataContext := c.MustGet("Db").(GameDLInterface)
	gameId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		err := dataContext.DeleteGame(gameId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "successful"})
		}
	}

}
