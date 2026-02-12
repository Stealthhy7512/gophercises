package cmd

import (
	"context"
	"fmt"

	"github.com/Stealthhy7512/gophercises/urlshort/repository"
	"github.com/Stealthhy7512/gophercises/urlshort/service"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var shortURL string

var shortenCmd = &cobra.Command{
	Use:   "shorten [url]",
	Short: "Shorten a URL.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := utils.LoadMongoConfig()

		client, err := mongo.Connect(options.Client().ApplyURI(cfg.URI))
		if err != nil {
			return err
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
			return err
		}

		s := service.NewURLService(repo)

		c := context.Background()

		longURL := args[0]
		shortURL, err := s.ShortenURL(c, longURL)
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), shortURL)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
