package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os" // NEW: ADDED TO ALTER WORKING DIRECTORY FOR BANNER LOADING PATHS
	"strings"
	"testing"
)

// NEW: TEST TO VERIFY THE FILE EXPORT FUNCTIONALITY AND HTTP HEADERS WITH PATH RECOVERY
func TestAsciiArtHandler_Export(t *testing.T) {
	// NEW: SAVE THE ORIGINAL DIRECTORY PATH BEFORE CHANGING IT
	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	// NEW: MOVE TO ROOT DIRECTORY (TWO LEVELS UP) SO BANNER LOADER CAN LOCATE THE .TXT FILES
	err = os.Chdir("../../")
	if err != nil {
		t.Fatalf("Failed to change directory to root: %v", err)
	}

	// NEW: ENSURE WE RETURN TO THE ORIGINAL TEST WORKING DIRECTORY AT THE END OF THE TEST
	defer func() {
		_ = os.Chdir(oldWD)
	}()

	// 1. CHOOSE A CONTROL TEXT AND BANNER FOR TESTING
	form := url.Values{}
	form.Set("text", "Hello")
	form.Set("banner", "standard")
	form.Set("action", "export") // CRITICAL: TRIGGERS THE EXPORT BLOCK IN YOUR HANDLER

	// 2. CREATE A NEW POST REQUEST WITH FORM DATA
	req, err := http.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// NEW: SET THE CORRECT CONTENT TYPE FOR FORM SUBMISSIONS
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 3. CREATE A RESPONSE RECORDER TO CAPTURE THE HANDLER'S OUTPUT
	rr := httptest.NewRecorder()

	// 4. CALL THE HANDLER DIRECTLY
	AsciiArtHandler(rr, req)

	// 5. TEST STATUS CODE
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// 6. NEW: TEST MANDATORY ZONE01 EXPORT HEADERS
	expectedContentType := "text/plain; charset=utf-8"
	if got := rr.Header().Get("Content-Type"); got != expectedContentType {
		t.Errorf("Expected Content-Type %q, got %q", expectedContentType, got)
	}

	expectedDisposition := `attachment; filename="ascii-art.txt"`
	if got := rr.Header().Get("Content-Disposition"); got != expectedDisposition {
		t.Errorf("Expected Content-Disposition %q, got %q", expectedDisposition, got)
	}

	// NEW: VERIFY CONTENT-LENGTH IS PRESENT AND NOT ZERO
	contentLength := rr.Header().Get("Content-Length")
	if contentLength == "" || contentLength == "0" {
		t.Error("Expected a valid non-zero Content-Length header, got empty or zero")
	}

	// 7. NEW: VERIFY THAT THE DOWNLOADED BODY IS NOT EMPTY AND CONTAINS ART
	bodyStr := rr.Body.String()
	if len(bodyStr) == 0 {
		t.Error("Expected downloaded file content to contain ASCII art, but body was empty")
	}

	// OPTIONAL CHECK: ENSURE IT DOES NOT CONTAIN THE HTML TEMPLATE BY ACCIDENT
	if strings.Contains(bodyStr, "<!DOCTYPE html>") {
		t.Error("Handler returned the HTML page instead of raw text file data")
	}
}
