package main

import (
	"fmt"

	"github.com/chrisevett/stool/harness"
	"github.com/gin-gonic/gin"
)

func main() {
	reader := FileConfigReader{}
	config, err := LoadConfig("./config.yml", reader)
	if err != nil {
		fmt.Println("Error calling LoadConfig")
		fmt.Println(err)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	if config.Verifier == "inspec" {
		router.Use(InspecMiddleware(config))
	}
	Routes(router)
	router.Run(":8080")
}

func InspecMiddleware(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		run := &harness.InspecRunner{}
		c.Set("runner", run)
		c.Set("path", config.TestPath)
		c.Next()
	}
}
