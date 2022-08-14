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

	// if _, err := conn.Exec(context.Background(), "INSERT INTO TODOS(title) VALUES($1)", "Hello db"); err != nil {
	// 	// Handling error, if occur
	// 	fmt.Println("Unable to insert due to: ", err)
	// 	return
	// }
	// fmt.Println("Insertion Succesfull")

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/todos", todoNewTaskHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

type todoNewTask struct {
	Title string `json:"title"`
}

func todoNewTaskHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t todoNewTask
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
