package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/Stealthhy7512/gophercises/urlshort/cmd"
	"github.com/lmittmann/tint"
)

func main() {
	var w = os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	cmd.Execute()
}
