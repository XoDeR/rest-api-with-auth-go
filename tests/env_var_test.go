package tests

import (
	"api-auth/config"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestEnvVar(t *testing.T) {
	// Change working directory to project root
	err := os.Chdir("../") // Moves up one level to project root
	if err != nil {
		log.Fatal("Failed to change directory:", err)
	}

	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	config.LoadEnvVar()
	dns := os.Getenv("DNS")
	t.Log(dns)
	t.Log("Running TestEnvVar")
}
