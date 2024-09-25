package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const (
	pingAddr            = "http://ping-pong-svc:3456"
	informationFilePath = "config/information.txt"
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

	pongResponse, err := http.Get(pingAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer pongResponse.Body.Close()

	body, err := io.ReadAll(pongResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	pongs := string(body)

	hash := uuid.New()

	information, err := os.ReadFile(informationFilePath)
	if err != nil {
		log.Fatalf("error when opening config/instructions.txt, %v", err)
	}

	envMessage := os.Getenv("MESSAGE")

	site := fmt.Sprintf("file content: %senv variable: MESSAGE=%s\n%s %s\nPing / Pongs: %s", information, envMessage, timestamp, hash, pongs)

	w.Write([]byte(site))
}
