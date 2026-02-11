package cmd

import (
	"maps"
	"net/http"
	"os"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/router"
	"github.com/spf13/cobra"
)

var (
	jsonPath string = ""
	yamlPath string = ""
)

var rootCmd = &cobra.Command{
	Use:   "urlshort",
	Short: "A shortened URL redirection service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var configs []map[string]string

		if jsonPath != "" {
			data, err := os.ReadFile(jsonPath)
			if err != nil {
				return err
			}

			cfg, err := handler.JSONHandler(data)
			if err != nil {
				return err
			}
			configs = append(configs, cfg.PathsToUrls)
		}

		if yamlPath != "" {
			data, err := os.ReadFile(yamlPath)
			if err != nil {
				return err
			}

			cfg, err := handler.YAMLHandler(data)
			if err != nil {
				return err
			}
			configs = append(configs, cfg.PathsToUrls)
		}

		finalConfig := mergeMaps(configs...)

		r := router.SetupRouter(&handler.MapHandler{
			PathsToUrls: finalConfig,
		})

		return http.ListenAndServe("localhost:8080", r)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&jsonPath, "json-path", "j", "", "Path to JSON config")
	rootCmd.Flags().StringVarP(&yamlPath, "yaml-path", "y", "", "Path to YAML config")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func mergeMaps(mapsToMerge ...map[string]string) map[string]string {
	merged := make(map[string]string)

	for _, m := range mapsToMerge {
		maps.Copy(merged, m)
	}

	return merged
}
