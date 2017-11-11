package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {

	err := verifyEnvironment()
	if err != nil {
		fmt.Println("Error invoking inspec. Please install the latest version of inspec from https://downloads.chef.io/inspec")
		return
	}
	router := gin.Default()
	router.Use(gin.Recovery())
	Routes(router)

	router.Run(":8080")
}

func verifyEnvironment() error {
	cmd := exec.Command("inspec", "version")
	err := cmd.Run()
	return err
}
