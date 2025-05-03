package tests

import (
	"api-auth/config"
	"api-auth/db/initializers"
	"api-auth/internal/models"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestDbPopulateWithModels(t *testing.T) {
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

	initializers.ConnectDB()

	// Drop all the tables
	err = initializers.DB.Migrator().DropTable(models.User{}, models.Tag{}, models.TimeSpan{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	// Migrate again
	err = initializers.DB.AutoMigrate(models.User{}, models.Tag{}, models.TimeSpan{})

	if err != nil {
		log.Fatal("Migration failed")
	}

	// Create sample data
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	tag := models.Tag{Color: "blue", Key: "dev", User: user}
	timeSpan := models.TimeSpan{Start: time.Now().Add(-24 * time.Hour), End: time.Now(), User: user, Tags: []models.Tag{tag}}

	// Insert data into the database
	err = initializers.DB.Create(&user).Error
	if err != nil {
		t.Fatal("Failed to create user:", err)
	}

	err = initializers.DB.Create(&tag).Error
	if err != nil {
		t.Fatal("Failed to create tag:", err)
	}

	err = initializers.DB.Create(&timeSpan).Error
	if err != nil {
		t.Fatal("Failed to create time span:", err)
	}

	t.Log("Successfully populated the database with test models")

	t.Log("Running TestDbPopulateWithModels")
}
