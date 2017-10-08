package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveAllPlayers(c *gin.Context) {
	dataContext := c.MustGet("Db").(PlayerDLInterface)
	player, err := dataContext.RetrieveAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": player})
	}
}

func RetrieveSinglePlayer(c *gin.Context) {
	dataContext := c.MustGet("Db").(PlayerDLInterface)
	gameId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		game := Player{Id: gameId}
		g, err := dataContext.RetrieveSinglePlayer(game.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": g})
		}
	}
}
