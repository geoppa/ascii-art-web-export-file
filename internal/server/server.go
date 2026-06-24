package server

import (
	"net/http"

	"ascii-art-web-export-file/internal/handlers"
)

// function tha creates a router
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	return mux
}

func Start() error {
	return http.ListenAndServe(":8080", NewRouter())
}
