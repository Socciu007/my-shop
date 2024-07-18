package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBService represents a service that interacts with a database.
type MongoDBService interface {
	// Health returns a map of health status information.
	Health() map[string]string // The keys and values in the map are service-specific.

	// Close terminates the database connection.
	Close() error // It returns an error if the connection cannot be closed.
}

// mongoService struct holds the database connection
type ClientType struct {
	client *mongo.Client
}

func NewMongoDBService(client *mongo.Client) *ClientType {
	return &ClientType{client: client}
}

// Health checks the health of the MongoDB connection.
func (ms *ClientType) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping MongoDB
	err := ms.client.Ping(ctx, nil)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("MongoDB connection error: %v", err)
		log.Fatalf("MongoDB connection error: %v", err) // Log the error and terminate the program
		return stats
	}

	// MongoDB is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "MongoDB is healthy"

	return stats
}

// Close terminates the MongoDB connection.
func (ms *ClientType) Close() error {
	err := ms.client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
		return err
	}
	log.Println("Disconnected from MongoDB")
	return nil
}
