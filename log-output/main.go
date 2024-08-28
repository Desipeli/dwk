package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

var logMessage string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	randomString := uuid.New()

	go logTimestampAndMessage(randomString.String())

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleGet)

	portAddr := ":" + port
	http.ListenAndServe(portAddr, mux)
}

func logTimestampAndMessage(msg string) {
	for {
		timestamp := time.Now().Format("2006-01-02 15:04:05.999 -0700")
		logMessage = fmt.Sprintf("%s: %s", timestamp, msg)
		fmt.Println(logMessage)
		time.Sleep(5 * time.Second)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(logMessage))
}
