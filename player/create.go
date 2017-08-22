package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(c *gin.Context) {
	var p Player
	if c.BindJSON(&p) == nil {
		if err := p.CreatePlayer(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": p.Id, "Name": p.Name})
		}
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
}
