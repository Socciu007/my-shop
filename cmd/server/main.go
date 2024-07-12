package main

import (
	"log"

	"my_shop/internal/config"
	"my_shop/internal/models"
	routers "my_shop/internal/routers"
	// "github.com/gin-gonic/gin"
)

func main() {
	// Switch to "release" mode in production.
	// gin.SetMode(gin.ReleaseMode)

	// Initialize MySQL service
	mysqlService := config.New()

	// Initialize MongoDB service
	mongoService, err := config.NewMongoDBService()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB service: %v", err)
	}

	// Perform helth check
	sqlHealth := mysqlService.Health()
	log.Printf("MySQL connection established")
	for key, value := range sqlHealth {
		log.Printf("%s: %s\n", key, value)
	}

	// Initialize and migrate models
	err = models.InitializeDB(mysqlService.db())
	if err != nil {
		panic(err)
	}

	mongoHealth := mongoService.Health()
	log.Printf("MongoDB connection established")
	for key, value := range mongoHealth {
		log.Printf("%s: %s\n", key, value)
	}

	// Close connections when done
	defer func() {
		if err := mysqlService.Close(); err != nil {
			log.Fatalf("Failed to close MySQL connection: %v", err)
		}

		if err := mongoService.Close(); err != nil {
			log.Fatalf("Failed to close MongoDB connection: %v", err)
		}
	}()

	router := routers.SetupRouter()
	router.Run(":8080")
}
