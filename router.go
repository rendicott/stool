package main

import (
	"github.com/gapi/butt"
	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	playerRouter := router.Group("/players")
	{
		playerRouter.POST("/", player.CreatePlayer)
		playerRouter.GET("/", player.PlayerIndex)
		playerRouter.GET("/:Id", player.ShowPlayer)
		playerRouter.DELETE("/:Id", player.DeletePlayer)
	}
	gameRouter := router.Group("/games")
	{
		gameRouter.POST("/", game.CreateGame)
		gameRouter.GET("/", game.GameIndex)
		gameRouter.GET("/:Id", game.ShowGame)
		gameRouter.DELETE("/:Id", game.DeleteGame)
	}
	outcomeRouter := router.Group("/outcomes")
	{
		outcomeRouter.POST("/", outcome.CreateOutcome)
		outcomeRouter.GET("/", outcome.OutcomeIndex)
		outcomeRouter.GET("/:Id", outcome.ShowOutcome)
		outcomeRouter.DELETE("/:Id", outcome.DeleteOutcome)
	}
	buttRouter := router.Group("/butts")
	{
		buttRouter.GET("/", butt.RetrieveAllButts)
		buttRouter.GET("/:Id", butt.RetrieveSingleButt)
	}
	router.Use(ButtDataContextMW())

}

func ButtDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		buttDl := &butt.ButtDLGorm{}
		c.Set("Db", buttDl)
		c.Next()
	}
}
