package cmd

import (
	"context"
	"net/http"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/repository"
	"github.com/Stealthhy7512/gophercises/urlshort/router"
	"github.com/Stealthhy7512/gophercises/urlshort/service"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the URL redirection server.",
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
		h := &handler.URLHandler{
			URLService: s,
		}

		r := router.SetupRouter(h)
		logger.Info("Starting server on port " + port)

		return http.ListenAndServe("localhost:"+port, r)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")

	rootCmd.AddCommand(serveCmd)
}
