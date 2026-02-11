package handler

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

type unmarshalFunc func([]byte, any) error

func QueryHandler(queryBytes []byte, fn unmarshalFunc) (*MapHandler, error) {
	paths, err := parseQuery(queryBytes, fn)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(paths)

	return &MapHandler{PathsToUrls: pathMap}, nil
}

func parseQuery(queryBytes []byte, fn unmarshalFunc) ([]parsedQuery, error) {
	pathUrls := []parsedQuery{}

	err := fn(queryBytes, &pathUrls)
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
