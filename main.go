package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/myapp")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	http.HandleFunc("/hello", helloHandler)
	http.Handle("/todos", todoHandler{
		conn: conn,
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

type todoNewTask struct {
	Title string `json:"title"`
}

type todoHandler struct {
	conn *pgx.Conn
}

func (h todoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t todoNewTask
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.conn.Exec(context.Background(), "INSERT INTO TODOS(title) VALUES($1)", t.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(map[string]string{"massage": "success"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
