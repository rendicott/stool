package main

import (
	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	playerRouter := router.Group("/player")
	{
		playerRouter.POST("/", player.CreatePlayer)
		playerRouter.GET("/", player.PlayerIndex)
		playerRouter.GET("/:Id", player.ShowPlayer)
		playerRouter.DELETE("/:Id", player.DeletePlayer)
	}
	gameRouter := router.Group("/game")
	{
		gameRouter.POST("/", game.CreateGame)
		gameRouter.GET("/", game.GameIndex)
		gameRouter.GET("/:Id", game.ShowGame)
		gameRouter.DELETE("/:Id", game.DeleteGame)
	}
	outcomeRouter := router.Group("/outcome")
	{
		outcomeRouter.POST("/", outcome.CreateOutcome)
		outcomeRouter.GET("/", outcome.OutcomeIndex)
		outcomeRouter.GET("/:Id", outcome.ShowOutcome)
		outcomeRouter.DELETE("/:Id", outcome.DeleteOutcome)
	}

}
