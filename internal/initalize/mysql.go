package initalize

import (
	"fmt"
	"log"
	"my_shop/internal/repo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *repo.DatabaseType{
	mysqlConfig := config.Mysql
	// fmt.Printf("%+v\n", mysqlConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	// Set connection pool settings
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
    }
    sqlDB.SetConnMaxLifetime(0)
    sqlDB.SetMaxIdleConns(50)
    sqlDB.SetMaxOpenConns(50)

	sqlService := repo.NewSQLService(db)
	log.Printf("MySQL connection established")

	// Health check
	statsSqlService := sqlService.Health()
	for key, value := range statsSqlService {
		log.Printf("%s: %s\n", key, value)
	}

	return sqlService
}