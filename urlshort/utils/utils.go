package utils

import (
	"encoding/json"
	"maps"
	"net/http"
)

type JsonResponse map[string]string

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func MergeMaps(mapsToMerge ...map[string]string) map[string]string {
	merged := make(map[string]string)

	for _, m := range mapsToMerge {
		maps.Copy(merged, m)
	}

	return merged
}
