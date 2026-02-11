package main

import (
	"net/http"

	"github.com/Stealthhy7512/gophercises/urlshort/handler"
	"github.com/Stealthhy7512/gophercises/urlshort/router"
)

const yaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution 
`

func main() {
	provider, err := handler.YAMLHandler([]byte(yaml))
	if err != nil {
		panic(err)
	}
	r := router.SetupRouter(provider)

	http.ListenAndServe("localhost:8080", r)
}
