package utils

import (
	"encoding/json"
	"maps"
	"net/http"
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
