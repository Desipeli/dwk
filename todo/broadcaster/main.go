package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	port := os.Getenv("PORT")

	if natsURL == "" {
		log.Fatal("provide env NATS_URL")
	}

	env := os.Getenv("ENV")
	log.Printf("ENVIRONMENT %s", env)
	log.Printf("Testing deployment print")

	var discordWebhookUrl string

	if env == "production" {
		discordWebhookUrl = os.Getenv("DISCORD_WEBHOOK_URL")
		if discordWebhookUrl == "" {
			log.Fatal("provide env DISCORD_WEBHOOK_URL")
		}
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("new_todo", func(m *nats.Msg) {
		content := fmt.Sprintf(`New todo created: "%s"`, string(m.Data))
		log.Printf("Received message: %s\n", content)
		msg := Message{Content: content}
		err := PostMessage(discordWebhookUrl, &msg)
		if err != nil {
			log.Println(err)
		}
	})

	nc.Subscribe("todo_done", func(m *nats.Msg) {
		content := fmt.Sprintf(`Todo done: "%s"`, string(m.Data))
		log.Printf("Done todo: %s\n", content)
		msg := Message{Content: content}
		err := PostMessage(discordWebhookUrl, &msg)
		if err != nil {
			log.Println(err)
		}
	})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", healthCheckHandler)
	mux.HandleFunc("GET /healthz", healthCheckHandler)

	var portAddr string
	if port != "" {
		portAddr = ":" + port
	} else {
		portAddr = ":8006"
	}

	log.Printf("listening port %s", portAddr)
	err = http.ListenAndServe(portAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func PostMessage(url string, message *Message) error {
	if url == "" { // Not in production
		log.Printf("message: %s", message.Content)
		return nil
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	return nil
}

type Message struct {
	Content string `json:"content"`
}
