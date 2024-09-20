package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

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

	content, err := os.ReadFile("files/shared/pings.txt")
	if err != nil {
		content = []byte("0")
	}
	pings, err := strconv.Atoi(string(content))

	if err != nil {
		w.Write([]byte("error when converting string to int: "))
		w.Write([]byte(err.Error()))
		return
	}

	pings++
	response := fmt.Sprintf("pong %d", pings)

	err = os.WriteFile("files/shared/pings.txt", []byte(string(pings)), 0644)
	if err != nil {
		log.Printf("error when writing pings to file: %v", err)
		w.Write([]byte("error when writing pings to file: "))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(response))
}
