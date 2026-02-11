package handler

import (
	"io"
	"os"
)

type parsedQuery struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

type RedirectProvider interface {
	GetURL(path string) (string, bool)
}

type MapHandler struct {
	PathsToUrls map[string]string
}

func (m *MapHandler) GetURL(path string) (string, bool) {
	url, ok := m.PathsToUrls[path]
	return url, ok
}

type decoderFunc func(r io.Reader, v any) error

func QueryHandler(f *os.File, fn decoderFunc) (*MapHandler, error) {
	pathUrls := []parsedQuery{}

	if err := fn(f, &pathUrls); err != nil {
		return nil, err
	}

	pathMap := buildMap(pathUrls)

	return &MapHandler{PathsToUrls: pathMap}, nil
}

func buildMap(parsedQueries []parsedQuery) map[string]string {
	paths := make(map[string]string)
	for _, p := range parsedQueries {
		paths[p.Path] = p.URL
	}

	return paths
}
