package main

import "github.com/gin-gonic/gin"
import "github.com/chrisevett/stool/harness"

func Routes(router *gin.Engine) {
	playerRouter := router.Group("/test")
	{
		playerRouter.GET("/", harness.RunAllTests)
	}

	router.Use(InspecMiddleware())
}

func InspecMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		runner := &harness.InspecRunner{}
		c.Set("runner", runner)
		c.Next()
	}
}
