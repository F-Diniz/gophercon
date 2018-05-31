package main

import (
	"log"
	"net/http"

	"github.com/F-Diniz/gophercon/pkg/routing"
	"os"
)

// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home
func main() {
	log.Printf("Service is starting...")

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("Service port wasn't set")
	}

	r := routing.BaseRouter()
	log.Fatal(http.ListenAndServe(":" + port, r))
}
