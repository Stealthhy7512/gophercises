package handler

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
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

func YAMLHandler(yamlBytes []byte) (*MapHandler, error) {
	paths, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(paths)

	return &MapHandler{PathsToUrls: pathMap}, nil
}

func JSONHandler(jsonBytes []byte) (*MapHandler, error) {
	paths, err := parseJSON(jsonBytes)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(paths)

	return &MapHandler{PathsToUrls: pathMap}, nil
}

func parseYAML(yamlBytes []byte) ([]parsedQuery, error) {
	var pathUrls []parsedQuery

	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return []parsedQuery{}, err
	}

	return pathUrls, nil
}

func parseJSON(jsonBytes []byte) ([]parsedQuery, error) {
	var pathUrls []parsedQuery

	err := json.Unmarshal(jsonBytes, &pathUrls)
	if err != nil {
		return []parsedQuery{}, err
	}

	return pathUrls, nil
}

func buildMap(parsedQueries []parsedQuery) map[string]string {
	paths := make(map[string]string)
	for _, p := range parsedQueries {
		paths[p.Path] = p.URL
	}

	return paths
}
