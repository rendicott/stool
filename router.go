package main

import "github.com/gin-gonic/gin"
import "github.com/chrisevett/stool/harness"

func Routes(router *gin.Engine) {
	testRouter := router.Group("/test")
	{
		testRouter.GET("/", harness.RunAllTests)
	}
}
