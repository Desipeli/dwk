package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var todos []string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", handleTodos)

	portAddr := ":" + port
	err := http.ListenAndServe(portAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, hx-target, hx-request, hx-trigger, hx-current-url")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		log.Print("GET TODOS")
		handleGetTodos(w, r)
	case http.MethodPost:
		log.Print("NEW TODO")
		handlePostTodos(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	response := getTodoLi()
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

func handlePostTodos(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTodo := strings.TrimSpace(r.FormValue("todo"))
	if len(newTodo) > 140 || len(newTodo) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todos = append(todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html")

	response := getTodoLi()
	w.Write([]byte(response))
}
func getTodoLi() string {
	var todoList string
	for _, todo := range todos {
		todoList += fmt.Sprintf("<li>%s</li>", todo)
	}

	return todoList
}
