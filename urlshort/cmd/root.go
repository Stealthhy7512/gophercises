package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

var rootCmd = &cobra.Command{
	Use:   "urlshort [command]",
	Short: "A URL shortening and redirection service.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

}
