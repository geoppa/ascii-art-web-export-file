package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ascii-art-web-export-file/internal/server"
)

// TestHomeEndpointDirectly validates the Home handler route behavior within the root workflow context.
func TestHomeEndpointDirectly(t *testing.T) {
	// Initialize the centralized application router instance
	router := server.NewRouter()

	// Formulate a standard valid GET request pointing to the home domain
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to initialize request model: %v", err)
	}

	// Instantiate a response recorder buffer
	recorder := httptest.NewRecorder()

	// Route the request straight through the active system mux router configuration
	router.ServeHTTP(recorder, req)

	// Verify that the network response status evaluates cleanly to 200 OK
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK for home endpoint, got: %d", recorder.Code)
	}

	// Confirm that signature corporate branding tags are safely present inside the output stream
	bodyString := recorder.Body.String()
	if !strings.Contains(bodyString, "ASCII Art Web") {
		t.Errorf("Expected HTML response page layout components to contain dashboard titles.")
	}
}
