package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server started in port %s", port)
	portAddr := ":" + port
	http.ListenAndServe(portAddr, nil)
}
