package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var (
	pingAddr            = "http://ping-pong-svc"
	informationFilePath = "config/information.txt"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	pingServiceAddr := os.Getenv("PING_ADDR")
	if pingServiceAddr != "" {
		pingAddr = pingServiceAddr
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleGet)

	portAddr := ":" + port

	log.Printf("Listening on port %s", port)
	http.ListenAndServe(portAddr, mux)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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

	var data PongData

	err = json.NewDecoder(pongResponse.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	pongs := data.Pongs

	hash := uuid.New()

	information, err := os.ReadFile(informationFilePath)
	if err != nil {
		log.Fatalf("error when opening config/instructions.txt, %v", err)
	}

	envMessage := os.Getenv("MESSAGE")

	site := fmt.Sprintf("file content: %s<br>env variable: MESSAGE=%s<br>%s %s<br>Ping / Pongs: %d", information, envMessage, timestamp, hash, pongs)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(site))
}

type PongData struct {
	Pongs int64
}
