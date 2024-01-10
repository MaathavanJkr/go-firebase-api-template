// cmd/app/main.go
package main

import (
	"log"
	"strconv"

	"github.com/go-firebase-api-template/pkg/config"
	"github.com/go-firebase-api-template/pkg/database"
	"github.com/go-firebase-api-template/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Load configuration settings
	configFilePath := "config.yaml"
	if err := config.LoadConfig(configFilePath); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Access configuration settings
	serviceAccountKeyPath := config.AppConfig.FirebaseServiceAccountKeyPath
	port := config.AppConfig.Port
	environment := config.AppConfig.Environment

	// Connect to Database
	firestoreClient, err := database.InitFirebaseApp(serviceAccountKeyPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Define routes
	routes.SetupRoutes(r, firestoreClient)

	// Print the environment and port
	log.Printf("Running in %s mode", environment)
	log.Printf("Listening on port %d", port)

	// Run the server
	r.Run(":" + strconv.Itoa(port))
}
