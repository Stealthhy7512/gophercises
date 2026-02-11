package router

import (
	"net/http"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/go-chi/chi/v5"
)

type JsonResponse map[string]string

func SetupRouter(yh *handler.YAMLHandler) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Get("/", hello)

	yamlHandler, err := handler.YAMLHandler([]byte(yh))

	return yamlHandler, err
}

func hello(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, JsonResponse{"message": "Hello, world!"})
}
