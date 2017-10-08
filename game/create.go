package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGame(c *gin.Context) {
	dataContext := c.MustGet("Db").(GameDLInterface)
	gameName := c.Param("Name")
	if gameName != "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
	if game, err := dataContext.CreateGame(gameName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Id": game.Id, "Name": game.Name})
	}
}
