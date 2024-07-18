package main

import (
	"my_shop/internal/initalize"
)

func main() {
	initalize.Run()
	// // Switch to "release" mode in production.
	// gin.SetMode(gin.ReleaseMode)

	// // Initialize MySQL service
	// mysqlService := config.New()

	// // Perform helth check
	// sqlHealth := mysqlService.Health()
	// log.Printf("MySQL connection established")
	// for key, value := range sqlHealth {
	// 	log.Printf("%s: %s\n", key, value)
	// }

	// db := mysqlService.GetDB()
	// models.InitializeDB(db)

	// // Initialize MongoDB service
	// mongoService, err := config.NewMongoDBService()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize MongoDB service: %v", err)
	// }

	// // Perform health check of mongodb connection
	// mongoHealth := mongoService.Health()
	// log.Printf("MongoDB connection established")
	// for key, value := range mongoHealth {
	// 	log.Printf("%s: %s\n", key, value)
	// }

	// // Close connections when done
	// defer func() {
	// 	if err := mysqlService.Close(); err != nil {
	// 		log.Fatalf("Failed to close MySQL connection: %v", err)
	// 	}

	// 	if err := mongoService.Close(); err != nil {
	// 		log.Fatalf("Failed to close MongoDB connection: %v", err)
	// 	}
	// }()

	// router := routers.SetupRouter(db)
	// router.Run(":8080")
}
