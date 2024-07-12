package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBService represents a service that interacts with a database.
type MongoDBService interface {
	// Health returns a map of health status information.
	Health() map[string]string // The keys and values in the map are service-specific.

	// Close terminates the database connection.
	Close() error // It returns an error if the connection cannot be closed.
}

// mongoService struct holds the database connection
type mongoService struct {
	client *mongo.Client
}

var (
	mongoURI = os.Getenv("MONGO_URI")
	mongoClient *mongo.Client
)

func NewMongoDBService() (MongoDBService, error) {
	// Reuse MongoDB connection
	if mongoClient != nil {
		return &mongoService{client: mongoClient}, nil
	}

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	mongoClient = client // Store the MongoDB client for reuse

	return &mongoService{client: client}, nil
}

// Health checks the health of the MongoDB connection.
func (ms *mongoService) Health() map[string]string {
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
func (ms *mongoService) Close() error {
	err := ms.client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
		return err
	}
	log.Println("Disconnected from MongoDB")
	return nil
}
