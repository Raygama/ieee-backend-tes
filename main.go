package main

import (
	"log"

	"backend/config"
	"backend/docs"
	"backend/models"
	"backend/routes"
	"backend/utils"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Swagger Files API"
	docs.SwaggerInfo.Description = "This is a simple Files."
	docs.SwaggerInfo.Version = "1.0"
	envHost := utils.Getenv("ENV_HOST", "localhost:8080")
	docs.SwaggerInfo.Host = envHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Get the global database instance
	db := config.ConnectDatabase()

	// Create a new gin router with default middleware
	r := routes.SetupRouter(db)

	// Auto-migrate the file model to the database
	err = db.AutoMigrate(&models.File{})
	if err != nil {
		log.Fatal("Error migrating database")
	}

	r.Run()
}
