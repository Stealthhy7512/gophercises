package cmd

import (
	"context"

	"github.com/Stealthhy7512/gophercises/urlshort/repository"
	"github.com/Stealthhy7512/gophercises/urlshort/service"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var longURL string
var shortURL string

var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Shorten a URL.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := utils.LoadMongoConfig()

		client, err := mongo.Connect(options.Client().ApplyURI(cfg.URI))
		if err != nil {
			logger.Warn("Failed to connect to MongoDB", "error", err)
		}
		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				logger.Warn("Error disconnecting from MongoDB: ", "error", err)
			}
		}()

		db := client.Database(cfg.Database)
		col := db.Collection(cfg.Collection)
		repo, err := repository.NewMongoURLRepository(col)
		if err != nil {
			logger.Warn("Failed to create URL repository", "error", err)
		}

		s := service.NewURLService(repo)

		c := context.Background()

		shortURL, err := s.ShortenURL(c, longURL)
		if err != nil {
			logger.Warn("Failed to shorten URL", "error", err)
			return nil
		}

		logger.Info("Shortened URL", "shortURL", shortURL)
		return nil
	},
}

func init() {
	shortenCmd.Flags().StringVarP(&longURL, "url", "u", "", "The URL to shorten")
	shortenCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(shortenCmd)
}
