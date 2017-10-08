package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(c *gin.Context) {
	dataContext := c.MustGet("Db").(PlayerDLInterface)
	playerName := c.Param("Name")
	if playerName != "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
	if player, err := dataContext.CreatePlayer(playerName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Id": player.Id, "Name": player.Name})
	}
}
