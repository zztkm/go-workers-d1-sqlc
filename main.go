//go:build js && wasm

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-michi/michi"
	"github.com/syumai/workers"
	_ "github.com/syumai/workers/cloudflare/d1" // register driver

	"github.com/zztkm/workers-d1/gen/sqlc"
)

type todoHandler struct {
	db      *sql.DB
	querier *sqlc.Queries
}

func newTodoHandler(db *sql.DB) *todoHandler {
	return &todoHandler{db: db, querier: sqlc.New(db)}
}

func (h *todoHandler) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type createTodoRequest struct {
	Title string `json:"title"`
}

func (h *todoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.querier.CreateTodo(r.Context(), req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		fmt.Fprintf(w, "Failed to encode response: %v", err)
	}
}

func (h *todoHandler) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo, err := h.querier.GetTodo(r.Context(), int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		fmt.Fprintf(w, "Failed to encode response: %v", err)
	}
}

func (h *todoHandler) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.querier.ListTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		fmt.Fprintf(w, "Failed to encode response: %v", err)
	}

}

func main() {
	db, err := sql.Open("d1", "DB")
	if err != nil {
		panic(err)
	}

	h := newTodoHandler(db)

	r := michi.NewRouter()
	r.HandleFunc("/", h.index)
	r.HandleFunc("GET /todos", h.listTodos)
	r.HandleFunc("POST /todos", h.createTodo)
	r.Route("/todos", func(sub *michi.Router) {
		sub.HandleFunc("GET /{id}", h.getTodo)
	})
	http.Handle("/", r)
	workers.Serve(nil)
}
