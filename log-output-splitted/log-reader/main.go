package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleGet)

	portAddr := ":" + port

	log.Printf("Listening on port %s", port)
	http.ListenAndServe(portAddr, mux)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	timestamp, err := os.ReadFile("files/pod/logs.log")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not read log file"))
		return
	}

	pings, err := os.ReadFile(("files/shared/pings.txt"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not read the ping file"))
	}

	hash := uuid.New()

	message := fmt.Sprintf("%s %s\nPing / Pongs: %s", timestamp, hash, pings)

	w.Write([]byte(message))
}
