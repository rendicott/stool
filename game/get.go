package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GameIndex(c *gin.Context) {
	games, err := GetGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": games})
	}
}

func ShowGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		g := Game{Id: gameId}
		g, err = g.GetGame()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": g.Id, "Name": g.Name})
		}
	}
}
