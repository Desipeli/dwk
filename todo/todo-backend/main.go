package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

var databaseURL string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	databaseURL = os.Getenv("DATABASE_URL")

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", handleTodos)
	mux.HandleFunc("GET /", handleRoot)

	portAddr := ":" + port
	err := http.ListenAndServe(portAddr, LoggingMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
	response, err := getTodoLi(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

func handlePostTodos(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTodo := strings.TrimSpace(r.FormValue("todo"))
	if len(newTodo) > 140 || len(newTodo) < 1 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(r.Context())

	_, err = conn.Exec(r.Context(), "INSERT INTO todos(todo) VALUES ($1)", newTodo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html")

	response, err := getTodoLi(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(response))
}
func getTodoLi(r *http.Request) (string, error) {
	var todoList string

	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		return "", err
	}
	defer conn.Close(r.Context())

	var todos []Todo

	rows, err := conn.Query(r.Context(), "SELECT id, todo FROM todos")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Text); err != nil {
			return "", err
		}
		todos = append(todos, todo)
	}

	for _, todo := range todos {
		todoList += fmt.Sprintf("<li id=%d>%s</li>", todo.Id, todo.Text)
	}

	return todoList, nil
}

type Todo struct {
	Id   int
	Text string
}

func LoggingMiddleware(next http.Handler) http.Handler {
	// Interceptor is needed to access the satuscode
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
		}

		// Restore the body so it can be read again by the next handler
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		rw := &responseWriterInterceptor{w, http.StatusOK}

		// The end point
		next.ServeHTTP(rw, r)

		log.Printf(
			"%s %s %s %s %s %d %v",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			string(body),
			r.Proto,
			rw.status,
			time.Since(start),
		)
	})
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriterInterceptor) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
