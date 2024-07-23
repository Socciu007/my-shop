package initalize

import (
	"fmt"
	"log"
	"my_shop/global"
	"my_shop/internal/models"
	"my_shop/internal/routers"
)

func Run() {
	// load configuration
	LoadConfig()
	InitLogger()

	sqlService := InitMySQL()
	global.GetDB = sqlService.GetDB()
	models.InitializeDB(global.GetDB)

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

	port := global.Config.Server.Port
	r := routers.SetupRouter()
	r.Run(":" + fmt.Sprintf("%d", port))
}