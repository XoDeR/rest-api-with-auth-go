package main

import (
	"api-auth/config"
	"fmt"
	"os"
)

func init() {
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	config.LoadEnvVar()
}

func main() {
	fmt.Println("Api-auth proj initted")
}
