package main

import (
	"context"
	"log/slog"

	"github.com/Stealthhy7512/gophercises/urlshort/cmd"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	logger := slog.Default()

	mongoConfig := utils.LoadMongoConfig()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoConfig.URI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		logger.Warn("Failed to connect to MongoDB", "error", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			logger.Warn("Error disconnecting from MongoDB: ", "error", err)
		} else {
			logger.Info("Disconnected from MongoDB.")
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logger.Warn("Failed to ping MongoDB:", "error", err)
	}
	logger.Info("Successfully connected to MongoDB")

	cmd.Execute()
}
