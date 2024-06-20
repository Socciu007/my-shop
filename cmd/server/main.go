package main

import (
	"fmt"
	"log"
	"my_shop/internal/config"
	routers "my_shop/internal/routers"
	// "github.com/gin-gonic/gin"
)

func main() {
	// Đặt chế độ của Gin framework thành release mode
	// gin.SetMode(gin.ReleaseMode)

	// Tạo một instance của service sử dụng hàm New từ package config
	service := config.New()

	// check connect mysql
	healthStats := service.Health()
	for key, value := range healthStats {
		fmt.Printf("%s: %s\n", key, value)
	}

	defer func() {// close connect to mySql
		if err := service.Close(); err != nil {
			log.Fatalf("Error closing the database connection: %v", err)
		}
	}()

	router := routers.SetupRouter()
	router.Run(":8080")
}
