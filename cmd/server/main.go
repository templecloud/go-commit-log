package main

import (
	"log"

	"github.com/templecloud/go-commit-log/internal/server"
)

// Create a simple HTTPServer listening on 8080.
//
func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
