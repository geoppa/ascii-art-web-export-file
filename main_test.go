package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"ascii-art-web-export-file/internal/server"
)

// TestHTTPRoutes performs an integration test over the entire routing matrix
// ensuring endpoints respond with correct status codes and headers.
func TestHTTPRoutes(t *testing.T) {
	// Initialize the central application router configuration
	router := server.NewRouter()

	// Define table-driven cases mapping out standard operational routes
	tests := []struct {
		name           string
		method         string
		targetRoute    string
		postData       url.Values
		expectedStatus int
		checkHeader    func(t *testing.T, resp *httptest.ResponseRecorder)
	}{
		{
			name:           "GET Home Domain Endpoint",
			method:         http.MethodGet,
			targetRoute:    "/",
			postData:       nil,
			expectedStatus: http.StatusOK,
			checkHeader:    nil,
		},
		{
			name:           "GET Unregistered Route - 404 Validation",
			method:         http.MethodGet,
			targetRoute:    "/invalid-path-target",
			postData:       nil,
			expectedStatus: http.StatusNotFound,
			checkHeader:    nil,
		},
		{
			name:           "POST Route Without Token - 405 Method Not Allowed",
			method:         http.MethodPost,
			targetRoute:    "/",
			postData:       nil,
			expectedStatus: http.StatusMethodNotAllowed,
			checkHeader:    nil,
		},
		{
			name:        "POST Submit Valid Form - Generate Action",
			method:      http.MethodPost,
			targetRoute: "/ascii-art",
			postData: url.Values{
				"text":   {"Hello"},
				"banner": {"standard"},
				"action": {"generate"},
			},
			expectedStatus: http.StatusOK,
			checkHeader:    nil,
		},
		{
			name:        "POST Export Action - TXT Format Headers Verification",
			method:      http.MethodPost,
			targetRoute: "/ascii-art",
			postData: url.Values{
				"text":   {"Hello"},
				"banner": {"standard"},
				"action": {"export"},
				"format": {"txt"},
			},
			expectedStatus: http.StatusOK,
			checkHeader: func(t *testing.T, resp *httptest.ResponseRecorder) {
				// Verify file transfer attachment instructions are present
				disposition := resp.Header().Get("Content-Disposition")
				if !strings.Contains(disposition, "attachment;") || !strings.Contains(disposition, ".txt") {
					t.Errorf("Expected valid TXT attachment headers, got: %s", disposition)
				}
			},
		},
		{
			name:        "POST Export Action - JSON Format Headers Verification",
			method:      http.MethodPost,
			targetRoute: "/ascii-art",
			postData: url.Values{
				"text":   {"Hello"},
				"banner": {"standard"},
				"action": {"export"},
				"format": {"json"},
			},
			expectedStatus: http.StatusOK,
			checkHeader: func(t *testing.T, resp *httptest.ResponseRecorder) {
				// Verify that response stream outputs clean JSON headers
				contentType := resp.Header().Get("Content-Type")
				if !strings.Contains(contentType, "application/json") {
					t.Errorf("Expected application/json header configuration, got: %s", contentType)
				}
			},
		},
	}

	// Loop through each operational server test scenario
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			// Formulate request structure depending on GET or POST payloads
			if tt.method == http.MethodPost && tt.postData != nil {
				req, err = http.NewRequest(tt.method, tt.targetRoute, strings.NewReader(tt.postData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(tt.method, tt.targetRoute, nil)
			}

			if err != nil {
				t.Fatalf("Failed to initialize system request model: %v", err)
			}

			// Instantiate an isolated response recorder to catch execution feedback
			recorder := httptest.NewRecorder()

			// Route request directly into our active router engine mux
			router.ServeHTTP(recorder, req)

			// Assert network statuses match target conditions
			if recorder.Code != tt.expectedStatus {
				t.Errorf("Route [%s] execution status failure. Expected %d, got %d",
					tt.targetRoute, tt.expectedStatus, recorder.Code)
			}

			// Execute conditional runtime header validations if declared
			if tt.checkHeader != nil {
				tt.checkHeader(t, recorder)
			}
		})
	}
}
