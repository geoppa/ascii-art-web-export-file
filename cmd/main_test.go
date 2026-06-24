package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"ascii-art-web-export-file/internal/server"
)

func TestMainIntegration(t *testing.T) {
	// Αλλάζουμε προσωρινά το working directory στο root του project
	// ώστε τα templates και τα banner files να εντοπίζονται σωστά.
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	// Δημιουργούμε έναν προσωρινό HTTP server για το test
	// χωρίς να ανοίξουμε πραγματικό port στο σύστημα.
	testServer := httptest.NewServer(server.NewRouter())
	defer testServer.Close()

	// Στέλνουμε GET request στην αρχική σελίδα.
	resp, err := http.Get(testServer.URL + "/")
	if err != nil {
		t.Fatalf("Αποτυχία σύνδεσης με τον server: %v", err)
	}
	defer resp.Body.Close()

	// Επιβεβαιώνουμε ότι η αρχική σελίδα επιστρέφει 200 OK.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Αναμενόμενο Status 200 OK, αλλά πήραμε: %d", resp.StatusCode)
	}

	// Στέλνουμε request σε route που δεν υπάρχει
	// και επιβεβαιώνουμε ότι επιστρέφεται 404 Not Found.
	resp404, err := http.Get(testServer.URL + "/non-existent-page-xyz")
	if err != nil {
		t.Fatalf("Αποτυχία σύνδεσης με τον server στο test 404: %v", err)
	}
	defer resp404.Body.Close()

	if resp404.StatusCode != http.StatusNotFound {
		t.Errorf("Αναμενόμενο Status 404 Not Found, αλλά πήραμε: %d", resp404.StatusCode)
	}
}
