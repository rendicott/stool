package main

import (
	"net/http"

	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.StaticFS("/frontend", http.Dir("./frontend"))
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
		gameRouter.GET("/", game.RetrieveAllGames)
		gameRouter.GET("/:Id", game.RetrieveSingleGame)
		gameRouter.DELETE("/:Id", game.DeleteGame)
	}
	outcomeRouter := router.Group("/outcomes")
	{
		outcomeRouter.POST("/", outcome.CreateOutcome)
		outcomeRouter.GET("/", outcome.OutcomeIndex)
		outcomeRouter.GET("/:Id", outcome.ShowOutcome)
		outcomeRouter.DELETE("/:Id", outcome.DeleteOutcome)
	}
	router.Use(GameDataContextMW())
	router.Use(PlayerDataContextMW())

}

func GameDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		gameDl := &game.GameDLGorm{}
		c.Set("Db", gameDl)
		c.Next()
	}
}

func PlayerDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerDl := &player.PlayerDLGorm{}
		c.Set("Db", playerDl)
		c.Next()
	}
}
