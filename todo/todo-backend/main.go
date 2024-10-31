package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

var databaseURL string
var natsURL string

func main() {

	log.Printf("Testing deployment print all three apps")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	databaseURL = os.Getenv("DATABASE_URL")
	natsURL = os.Getenv("NATS_URL")

	mux := http.NewServeMux()
	mux.HandleFunc("/todos/{id}", handleTodoDone)
	mux.HandleFunc("/todos", handleTodos)
	mux.HandleFunc("/healthz", handleHealthCheck)
	mux.HandleFunc("/", handleRoot)

	portAddr := ":" + port
	err := http.ListenAndServe(portAddr, LoggingMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func handleTodoDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")

	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(r.Context())

	var todoText string

	err = conn.QueryRow(r.Context(), "UPDATE todos SET done=true WHERE id=$1 RETURNING todo", id).Scan(&todoText)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(todoText)

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()
	nc.Publish("todo_done", []byte(todoText))

	w.WriteHeader(http.StatusOK)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		conn, err := pgx.Connect(r.Context(), databaseURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer conn.Close(r.Context())
		w.WriteHeader(http.StatusOK)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
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

	todos, err := getTodos(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(todos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func handlePostTodos(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}

	var input TodoInput

	err = json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(input.Todo) > 140 || len(input.Todo) < 1 {
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

	var newTodo Todo
	err = conn.QueryRow(r.Context(), "INSERT INTO todos(todo) VALUES ($1) RETURNING id, todo, done", input.Todo).Scan(&newTodo.Id, &newTodo.Text, &newTodo.Done)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()
	nc.Publish("new_todo", []byte(newTodo.Text))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)

}
func getTodos(r *http.Request) ([]Todo, error) {
	conn, err := pgx.Connect(r.Context(), databaseURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close(r.Context())

	rows, err := conn.Query(r.Context(), "SELECT id, todo, done FROM todos WHERE done=false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Text, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

type Todo struct {
	Id   int
	Text string
	Done bool
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

type TodoInput struct {
	Todo string `json:"todo"`
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriterInterceptor) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func getRequestURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host
	return fmt.Sprintf("%s://%s%s", scheme, host, r.URL.RequestURI())
}
