package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		g := Game{Id: gameId}
		err = g.DeleteGame()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": g.Id, "Name": g.Name})
		}
	}
}
