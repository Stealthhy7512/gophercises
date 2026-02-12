package router

import (
	"net/http"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/utils"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(p handler.RedirectProvider) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", hello)

	r.Get("/{path}", func(w http.ResponseWriter, r *http.Request) {
		path := chi.URLParam(r, "path")

		if url, ok := p.GetURL(path); ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		utils.WriteJSON(w, http.StatusNotFound, utils.JsonResponse{"error": "Not found"})
	})

	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, utils.JsonResponse{"message": "Hello, world!"})
}
