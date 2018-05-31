package main

import (
	"log"
	"os"

	"github.com/F-Diniz/gophercon/pkg/routing"
	"github.com/F-Diniz/gophercon/webserver"
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
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())
}
