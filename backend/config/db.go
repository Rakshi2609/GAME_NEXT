package config

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectMongo() error {
	if mongoClient != nil {
		return nil
	}
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return errors.New("MONGO_URI not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	mongoClient = client
	return nil
}

func MongoClient() *mongo.Client { return mongoClient }

func DB() *mongo.Database {
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "codechef_onboarding"
	}
	return mongoClient.Database(name)
}
