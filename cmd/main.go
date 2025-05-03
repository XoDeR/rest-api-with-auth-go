package main

import (
	"api-auth/api/router"
	"api-auth/config"
	"api-auth/db/initializers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	config.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Api-auth proj initted")

	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}
