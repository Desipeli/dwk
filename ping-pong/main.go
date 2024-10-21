package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

var databaseURL string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	databaseURL = os.Getenv("DATABASE_URL")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pingpong", handleGetPing)
	mux.HandleFunc("GET /healthz", handleHealthCheck)
	mux.HandleFunc("GET /", handleGetCurrentPongs)

	portAddr := ":" + port

	log.Printf("Listening on port: %s", port)
	http.ListenAndServe(portAddr, mux)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		log.Printf("Health check database error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	conn.Close(r.Context())
	w.WriteHeader(http.StatusOK)
}

func handleGetCurrentPongs(w http.ResponseWriter, r *http.Request) {
	pongs, err := getPongsFromDb(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := map[string]int64{
		"pongs": pongs,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleGetPing(w http.ResponseWriter, r *http.Request) {

	pongs, err := increasePongs(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("pongs %d", pongs)

	w.Write([]byte(response))
}

func getPongsFromDb(r *http.Request) (int64, error) {

	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer conn.Close(r.Context())

	var pongs int64
	err = conn.QueryRow(r.Context(), "SELECT pongs FROM pings").Scan(&pongs)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return pongs, nil
}

func increasePongs(r *http.Request) (int64, error) {
	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer conn.Close(r.Context())

	_, err = conn.Exec(r.Context(), "UPDATE pings SET pongs = pongs+1 WHERE id=1")
	if err != nil {
		return 0, err
	}

	var pongs int64
	err = conn.QueryRow(r.Context(), "SELECT pongs FROM pings").Scan(&pongs)
	if err != nil {
		return 0, err
	}

	return pongs, nil
}
