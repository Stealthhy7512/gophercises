package handler

import (
	"context"

	"github.com/Stealthhy7512/gophercises/urlshort/service"
)

type RedirectProvider interface {
	GetURL(path string) (string, bool)
}

type URLHandler struct {
	URLService service.URLService
}

func (u *URLHandler) ShortenURL(c context.Context, longURL string) (string, error) {
	return u.URLService.ShortenURL(c, longURL)
}

// GetURL checks if the given path exists in the PathsToUrls map
// and returns the corresponding URL and a boolean indicating if it was found.
func (u *URLHandler) GetURL(path string) (string, bool) {
	shortURL := path

	longURL, err := u.URLService.GetLongURL(context.Background(), shortURL)
	if err != nil {
		return "", false
	}

	return longURL, true
}
