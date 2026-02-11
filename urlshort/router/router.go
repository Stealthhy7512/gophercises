package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type JsonResponse map[string]string

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", hello)

	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, JsonResponse{"message": "Hello, world!"})
}
