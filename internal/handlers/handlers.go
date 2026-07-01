package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"ascii-art-web-export-file/internal/banner"
	"ascii-art-web-export-file/internal/render"
	"ascii-art-web-export-file/internal/validation"
)

// Pre-parse templates to avoid repetitive and costly disk I/O operations per request
var (
	homeTemplate  *template.Template
	errorTemplate *template.Template
)

func init() {
	// template.Must will panic if the files are missing at startup,
	// enforcing a fail-fast strategy for deployment issues.
	homeTemplate = template.Must(template.ParseFiles("templates/index.html"))
	errorTemplate = template.Must(template.ParseFiles("templates/error.html"))
}

// PageData holds form and result information for the main UI
type PageData struct {
	Text   string
	Banner string
	Error  string
	Result string
}

// ErrorPageData holds values to display custom error screens
type ErrorPageData struct {
	Code    int
	Message string
}

// renderHomePage displays the primary user submission layout
func renderHomePage(w http.ResponseWriter, data PageData) {
	err := homeTemplate.Execute(w, data)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

// renderErrorPage outputs dedicated custom error HTML screens
func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	// Set the response status code header before writing any body content
	w.WriteHeader(statusCode)

	err := errorTemplate.Execute(w, ErrorPageData{
		Code:    statusCode,
		Message: message,
	})
	if err != nil {
		// Fallback to plain text if the template rendering fails entirely
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HomeHandler serves the dashboard interface and protects against broken paths
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Reject paths that do not exactly match the home domain root
	if r.URL.Path != "/" {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}

	// Reject any incoming method that is not a standard GET request
	if r.Method != http.MethodGet {
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	renderHomePage(w, PageData{})
}

// AsciiArtHandler validates form inputs, reads files, and processes the output art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Block requests trying to execute anything other than a secure POST
	if r.Method != http.MethodPost {
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Limit the request body size (e.g., 10MB) to protect against DoS attacks
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	// Read form inputs safely
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest, "Bad Request")
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")
	action := r.FormValue("action") // Captures which submit button was pressed (generate or export)

	// Verify inputted text conforms to printable character rules
	err = validation.ValidateText(text)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Tells the browser this is a 400 error
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	// Verify the chosen template target name matches existing files
	err = validation.ValidateBanner(bannerName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Force a 400 Bad Request network header
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	bannerFile := bannerName + ".txt"

	// Read text representation matrix out of storage
	bannerData, err := banner.Load(bannerFile)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Generate standard visual ASCII character block map
	result := render.Generate(text, bannerData)

	// Check if the user requested to download the file
	if action == "export" {
		format := r.FormValue("format")
		if format == "" {
			format = "txt"
		}

		var dataBytes []byte
		var fileName string
		var err error

		switch format {
		case "html":
			htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>ASCII Art - %s</title>
    <style>
        body {
            background-color: #02040f;
            color: #22d3ee;
            font-family: "Courier New", Courier, monospace;
            padding: 30px;
            white-space: pre;
            font-size: 14px;
            line-height: 1.2;
        }
    </style>
</head>
<body>%s</body>
</html>`, bannerName, result)

			dataBytes = []byte(htmlContent)
			fileName = fmt.Sprintf("ascii-art-%s.html", bannerName)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

		case "json":
			// Create a structured map to hold the data cleanly
			jsonData := map[string]string{
				"banner": bannerName,
				"input":  text,
				"ascii":  result,
			}

			// MarshalIndent makes the downloaded JSON file look pretty and readable
			dataBytes, err = json.MarshalIndent(jsonData, "", "  ")
			if err != nil {
				renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			fileName = fmt.Sprintf("ascii-art-%s.json", bannerName)
			w.Header().Set("Content-Type", "application/json")

		case "txt":
			fallthrough
		default:
			dataBytes = []byte(result)
			fileName = fmt.Sprintf("ascii-art-%s.txt", bannerName)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		}

		// Send correct download behavior instructions back to browser client
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		w.Header().Set("Content-Length", strconv.Itoa(len(dataBytes)))
		w.Write(dataBytes)
		return
	}

	// Send execution success back to UI output target elements
	renderHomePage(w, PageData{
		Text:   text,
		Banner: bannerName,
		Result: result,
	})
}
