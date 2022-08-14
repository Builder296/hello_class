package main

import (
	"bytes"
	"fmt"
	"testing"
)

type fakeStore struct{}

func (fakeStore) NewTask(title string) error { return nil }

func TestTodosNewTaskHandler(t *testing.T) {
	payload := bytes.NewBuffer([]byte`{"title": "Test New Task"}`)

	req := httptest.NewRequest(http.MethodPost, "/todos", payload)
    w := httptest.NewRecorder()
	h := todoHandler(store: &fakeStore{})
	h.ServeHTTP(w, req)

    resp := w.Result()
	// TODO
	
    defer resp.Body.Close()

	fmt.Println("")
}
