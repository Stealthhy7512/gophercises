package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var logger = slog.Default()

var rootCmd = &cobra.Command{
	Use:   "urlshort",
	Short: "A shortened URL redirection service.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

}
