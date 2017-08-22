package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGame(c *gin.Context) {
	var g Game
	if c.BindJSON(&g) == nil {
		if err := g.CreateGame(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": g.Id, "Name": g.Name})
		}
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
}
