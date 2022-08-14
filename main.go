package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/myapp")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	http.HandleFunc("/hello", helloHandler)

	todoHandler := todoHandler{store: NewStore(conn)}
	http.Handle("/todos", todoHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINFO)
	defer stop()

	srv := &http.Server{Addr: ":8081"}

	go func() {
		srv.ListenAndServe()
	}()

	<-ctx.Done()
	stop()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Println(err)
	}

	fmt.Println("Server stopped")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

type store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *store {
	return &store{conn: conn}
}

func (s store) NewTask(title string) error {
	if _, err := s.conn.Exec(context.Background(), "INSERT INTO TODOS(title) VALUES($1)", title); err != nil {
		return err
	}
	return nil
}

type todoNewTask struct {
	Title string `json:"title"`
}

type todoHandler struct {
	store *store
}

func (h todoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t todoNewTask
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.store.NewTask(t.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(map[string]string{"massage": "success"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
