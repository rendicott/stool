package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PlayerIndex(c *gin.Context) {
	players, err := GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": players})
	}
}

func ShowPlayer(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		p := Player{Id: playerId}
		p, err := p.GetPlayer()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": p.Id, "Name": p.Name})
		}
	}
}
