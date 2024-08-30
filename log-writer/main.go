package main

import (
	"log"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	for {
		logTimestampToFile("logs.log")
		time.Sleep(5 * time.Second)
	}

}

func logTimestampToFile(filename string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.999 -0700")
	err := os.WriteFile(filename, []byte(timestamp), 0644)
	if err != nil {
		log.Printf("error when writing to file: %v", err)
	}
}
