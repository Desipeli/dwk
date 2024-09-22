package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const pingsPath = "files/pings.txt"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pingpong", handleGetPing)
	mux.HandleFunc("GET /", handleGetCurrentPongs)

	portAddr := ":" + port

	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(portAddr, mux)
}

func handleGetCurrentPongs(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile(pingsPath)
	if err != nil {
		content = []byte("0")
	}

	w.Write([]byte(content))
}

func handleGetPing(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile(pingsPath)
	if err != nil {
		content = []byte("0")
	}
	pings, err := strconv.Atoi(string(content))

	if err != nil {
		pings = 0
	}

	pings++
	response := fmt.Sprintf("pong %d", pings)

	err = os.WriteFile(pingsPath, []byte(strconv.Itoa(pings)), 0644)
	if err != nil {
		log.Printf("error when writing pings to file: %v", err)
		w.Write([]byte("error when writing pings to file: "))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(response))
}
