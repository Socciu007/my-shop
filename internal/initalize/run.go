package initalize

import (
	"fmt"
	"log"
	"my_shop/internal/models"
	"my_shop/internal/routers"
)

func Run() {
	// load configuration
	LoadConfig()
	InitLogger()
	mysqlConfig := config
	fmt.Printf("%+v\n", mysqlConfig)

	sqlService := InitMySQL()
	getDb := sqlService.GetDB()
	models.InitializeDB(getDb)

	mongoService := InitMongoDB()

	// Close connections when done
	defer func() {
		if err := sqlService.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}

		if err := mongoService.Close(); err != nil {
			log.Fatalf("Failed to close MongoDB connection: %v", err)
		}
	}()

	r := routers.SetupRouter(getDb)
	r.Run(":8080")
}