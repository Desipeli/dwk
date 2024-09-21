package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"todo-app/internal/templates"
)

const (
	persistentPicTimestamp = "picDownloaded.txt"
	imageFile              = "./public/image.jpg"
	imageDuration          = 1 * time.Hour
	timeLayout             = "2006-01-02 15:04:05.999 -0700"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	mux.HandleFunc("/", homePage)

	log.Printf("Server started on port: %s", port)
	portAddr := ":" + port
	err := http.ListenAndServe(portAddr, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != ("/") {
		// catch all undefined routes
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := downloadImageIfTimePassed()
	if err != nil {
		log.Print(err)
	}

	page := templates.HomePage()
	page.Render(r.Context(), w)

}

func downloadImageIfTimePassed() error {
	content, err := os.ReadFile(persistentPicTimestamp)
	if err != nil {
		err = downloadNewImage()
		if err != nil {
			return err
		}
	} else {
		timestamp, err := time.Parse(timeLayout, string(content))
		if err != nil {
			return err
		}
		timePassedFromLastDownload := time.Since(timestamp)

		if timePassedFromLastDownload > imageDuration {
			err = downloadNewImage()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func downloadNewImage() error {
	log.Print("Start image download")
	url := "https://picsum.photos/1200"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(imageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format(timeLayout)

	err = os.WriteFile(persistentPicTimestamp, []byte(timestamp), 0644)
	if err != nil {
		return err
	}

	return nil

}
