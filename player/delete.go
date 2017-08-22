package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePlayer(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		p := Player{Id: playerId}
		err = p.DeletePlayer()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": p.Id, "Name": p.Name})
		}
	}
}
