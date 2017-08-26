package butt

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveAllButts(c *gin.Context) {
	dataContext := c.MustGet("Db").(ButtDLInterface)
	butts, err := dataContext.RetrieveAllButts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": butts})
	}
}

func RetrieveSingleButt(c *gin.Context) {
	dataContext := c.MustGet("Db").(ButtDLInterface)
	buttId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		butt := Butt{Id: buttId}
		b, err := dataContext.RetrieveSingleButt(butt.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Id": b.Id, "Name": b.Name})
		}
	}
}
