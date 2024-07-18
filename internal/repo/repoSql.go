package repo

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type ServiceSql interface {
	// Health returns a map of health status information.
	Health() map[string]string // The keys and values in the map are service-specific.

	// Close terminates the database connection.
	Close() error // It returns an error if the connection cannot be closed.

	GetDB()
}

// Service struct holds the database connection
type DatabaseType struct {
	db *gorm.DB
}

func NewSQLService(db *gorm.DB) *DatabaseType {
	return &DatabaseType{db: db}
}

// GetDB returns the database connection.
func (s *DatabaseType) GetDB() *gorm.DB {
	return s.db
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
// Health checks the health of the database connection by pinging the database.
func (s *DatabaseType) Health() map[string]string {
	sqlDB, err := s.db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err = sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("MySQL connection error: %v", err)
		log.Fatalf("MySQL connection error: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "MySQL is healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}
	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}
	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Closes the database connection.
func (s *DatabaseType) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("error getting DB instance: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %v", err)
	}

	return nil
}