package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Handle request
		requestPath := r.URL.Path

		if redirectUrl, ok := pathsToUrls[requestPath]; ok {
			// Write response
			http.Redirect(w, r, redirectUrl, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// Parse YAML
	pathUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	// Build map
	pathMap := buildMap(pathUrls)

	// Return map handler
	return MapHandler(pathMap, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(data []byte) ([]PathUrl, error) {

	var pathUrls []PathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return pathUrls, err
}

func buildMap(pathUrls []PathUrl) map[string]string {
	pathMap := make(map[string]string)
	for _, pu := range pathUrls {
		pathMap[pu.Path] = pu.URL
	}
	return pathMap
}
