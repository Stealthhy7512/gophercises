package utils

import (
	"encoding/json"
	"log/slog"
	"maps"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type JsonResponse map[string]string

// WriteJSON is a helper function to write a JSON response with the given status code and data.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// MergeMaps takes multiple maps and merges them into a single map.
// If there are duplicate keys, the value from the last map will be used.
func MergeMaps(mapsToMerge ...map[string]string) map[string]string {
	merged := make(map[string]string)

	for _, m := range mapsToMerge {
		maps.Copy(merged, m)
	}

	return merged
}

type MongoConfig struct {
	URI        string
	Database   string
	Collection string
}

func LoadMongoConfig() *MongoConfig {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, using environment variables.")
	}

	return &MongoConfig{
		URI:        getEnv("MONGO_URI"),
		Database:   getEnv("DATABASE_NAME"),
		Collection: getEnv("COLLECTION_NAME"),
	}
}

func getEnv(k string) string {
	if value, ok := os.LookupEnv(k); ok {
		return value
	}
	return ""
}
