package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	randomString := uuid.New()

	for {
		log.Print(randomString)
		time.Sleep(5 * time.Second)
	}
}
