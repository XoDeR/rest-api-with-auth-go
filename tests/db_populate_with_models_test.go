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
	// explicitly drop m2m timespan_tags because it's not dropped automatically
	err = initializers.DB.Migrator().DropTable("timespan_tags", models.User{}, models.Tag{}, models.TimeSpan{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	// Migrate again
	err = initializers.DB.AutoMigrate(models.User{}, models.Tag{}, models.TimeSpan{})

	if err != nil {
		log.Fatal("Migration failed")
	}

	// Create sample data
	// Insert data into the database
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	err = initializers.DB.Create(&user).Error
	if err != nil {
		t.Fatal("Failed to create user:", err)
	}

	tag := models.Tag{Color: "blue", Key: "dev", User: user}
	err = initializers.DB.Create(&tag).Error
	if err != nil {
		t.Fatal("Failed to create tag:", err)
	}

	timeSpan := models.TimeSpan{Start: time.Now().Add(-24 * time.Hour), End: time.Now(), User: user, Tags: []models.Tag{tag}}
	err = initializers.DB.Create(&timeSpan).Error
	if err != nil {
		t.Fatal("Failed to create time span:", err)
	}

	t.Log("Successfully populated the database with test models")

	t.Log("Running TestDbPopulateWithModels")
}
