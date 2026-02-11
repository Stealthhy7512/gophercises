package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/router"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	jsonPath string
	yamlPath string

	port string
)

var rootCmd = &cobra.Command{
	Use:   "urlshort",
	Short: "A shortened URL redirection service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var configs []map[string]string

		if jsonPath != "" {
			file, err := os.Open(jsonPath)
			if err != nil {
				return err
			}
			defer file.Close()

			cfg, err := handler.QueryHandler(file, func(r io.Reader, v any) error {
				return json.NewDecoder(r).Decode(v)
			})
			if err != nil {
				return err
			}

			configs = append(configs, cfg.PathsToUrls)
		}

		if yamlPath != "" {
			file, err := os.Open(yamlPath)
			if err != nil {
				return err
			}
			defer file.Close()

			cfg, err := handler.QueryHandler(file, func(r io.Reader, v any) error {
				return yaml.NewDecoder(r).Decode(v)
			})
			if err != nil {
				return err
			}

			configs = append(configs, cfg.PathsToUrls)
		}

		finalConfig := utils.MergeMaps(configs...)

		r := router.SetupRouter(&handler.MapHandler{
			PathsToUrls: finalConfig,
		})

		return http.ListenAndServe("localhost:"+port, r)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&jsonPath, "json-path", "j", "", "Path to JSON config")
	rootCmd.Flags().StringVarP(&yamlPath, "yaml-path", "y", "", "Path to YAML config")
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
