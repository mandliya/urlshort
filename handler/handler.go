package handler

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler : takes in a mapping of paths to URLs, and redirects using http.handler
func MapHandler(pathsToURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToURL[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pu struct {
	path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(data []byte) ([]pu, error) {
	var pathsToURL []pu
	err := yaml.Unmarshal(data, &pathsToURL)
	return pathsToURL, err
}

func buildMap(pathsToURL []pu) map[string]string {
	pathsToURLMap := make(map[string]string)
	for _, pu := range pathsToURL {
		pathsToURLMap[pu.path] = pu.URL
	}
	fmt.Printf("length: %d, map %v\n", len(pathsToURLMap), pathsToURLMap)
	return pathsToURLMap
}

func printMap(puMap map[string]string) {
	for key, val := range puMap {
		fmt.Println(key, val)
	}
}

// YAMLHandler : takes in a yaml of paths to URLs, and redirects using http.handler
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToURL, err := parseYAML(data)
	if err != nil {
		return nil, err
	}
	pathsToURLMap := buildMap(pathsToURL)
	printMap(pathsToURLMap)
	return MapHandler(pathsToURLMap, fallback), nil
}
