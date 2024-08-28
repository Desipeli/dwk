package main

import (
	"log"
	"net/http"
	"os"
	"todo-app/internal/templates"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homePage)

	log.Printf("Server started in port: %s", port)
	portAddr := ":" + port
	http.ListenAndServe(portAddr, mux)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != ("/") {
		// catch all undefined routes
		w.WriteHeader(http.StatusNotFound)
		return
	}
	page := templates.HomePage()
	page.Render(r.Context(), w)

}
