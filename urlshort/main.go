package main

import (
	"net/http"

	"github.com/Stealthhy7512/gophercises/urlshort/router"
)

func main() {
	r := router.SetupRouter()

	http.ListenAndServe("localhost:8080", r)
}
