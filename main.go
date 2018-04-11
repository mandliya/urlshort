package main

import (
	"fmt"
	"net/http"
	"urlshort/handler"
)

func main() {
	mux := defaultMux()
	pathToURLs := map[string]string{
		"/google":   "http://google.com",
		"/bing":     "http://bing.com",
		"/yahoo/in": "http://yahoo.in",
	}
	mapHandler := handler.MapHandler(pathToURLs, mux)

	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := handler.YAMLHandler([]byte(yml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Stating the server at 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
