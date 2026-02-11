package utils

import (
	"encoding/json"
	"net/http"
)

type JsonResponse map[string]string

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
