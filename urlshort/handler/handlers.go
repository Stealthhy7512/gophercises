package handler

import (
	"gopkg.in/yaml.v3"
)

type parsedYAML struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
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

func parseYAML(yamlBytes []byte) ([]parsedYAML, error) {
	var pathUrls []parsedYAML

	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return []parsedYAML{}, err
	}

	return pathUrls, nil
}

func buildMap(parsedYAMLs []parsedYAML) map[string]string {
	paths := make(map[string]string)
	for _, p := range parsedYAMLs {
		paths[p.Path] = p.URL
	}
	return paths
}
