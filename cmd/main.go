package main

import (
	"fmt"
	"log"

	"ascii-art-web-export-file/internal/server"
)

func main() {
	// print a message to show the server is starting
	fmt.Println("Server starting at http://localhost:8080")
	// start the web server and look for incoming requests
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
