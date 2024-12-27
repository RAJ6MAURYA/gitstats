package main

import (
	"log"
	"net/http"

	_ "github.com/sonichigo/hg/routes"

	cu "github.com/sonichigo/hg/common"
)

func main() {
	// Get port from environment variable or use default
	port := cu.GetPort()

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
