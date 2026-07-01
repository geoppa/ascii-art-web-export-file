package main

import (
	"testing"

	"ascii-art-web-export-file/internal/server"
)

// TestNewRouterInitialization verifies that the server package correctly
// initializes the HTTP ServeMux router and compiles the registered routes.
func TestNewRouterInitialization(t *testing.T) {
	// Invoke the router factory function directly from the server package
	mux := server.NewRouter()

	// Ensure the returned pointer is initialized and not nil
	if mux == nil {
		t.Fatalf("Expected NewRouter() to return a valid http.ServeMux instance, got nil")
	}
}
