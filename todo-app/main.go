package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("Server started in port %s", port)
}
