package initalize

import (
	"context"
	"log"
	"my_shop/internal/repo"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *repo.ClientType{
	uri := config.MongoDB.URI
	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	mongoService := repo.NewMongoDBService(client)
	log.Printf("MongoDB connection established")

	// Perform health check of mongodb connection
	mongoHealth := mongoService.Health()
	for key, value := range mongoHealth {
		log.Printf("%s: %s\n", key, value)
	}

	return mongoService
}