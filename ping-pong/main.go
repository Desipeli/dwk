package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter int

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pingpong", handleGetPingPong)

	portAddr := ":" + port

	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(portAddr, mux)
}

func handleGetPingPong(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("pong %d", counter)
	counter++

	err := os.WriteFile("files/shared/pings.txt", []byte(string(counter)), 0644)
	if err != nil {
		log.Printf("error when writing pings to file: %v", err)
	}

	w.Write([]byte(response))
}
