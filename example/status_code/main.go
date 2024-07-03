package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zhassymov/result"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	ErrID       = errors.New("invalid id")
	ErrNotFound = errors.New("entity not found")
)

type Entity struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}

type EntityRepository interface {
	GetByID(ctx context.Context, id int64) (Entity, error)
}

type memoryRepository struct {
	mu       sync.RWMutex
	entities map[int64]Entity
}

func (m *memoryRepository) GetByID(ctx context.Context, id int64) (Entity, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entity, ok := m.entities[id]
	if !ok {
		return Entity{}, ErrNotFound
	}

	return entity, nil
}

func HandleStatus(handle func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handle(w, r)
		if err == nil {
			return
		}

		// get status code from error
		status := result.StatusCode(err)

		// write error status code
		w.WriteHeader(status)

		// set error message
		msg := err.Error()

		// reset error message if error on server side
		if status >= http.StatusInternalServerError {
			msg = "something went wrong"
		}

		// write error message
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
	}
}

func GetEntityByID(entities EntityRepository) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			return result.BadRequest(ErrID)
		}

		entity, err := entities.GetByID(r.Context(), id)
		if errors.Is(err, ErrNotFound) {
			return result.NotFound(err)
		}
		if err != nil {
			return err
		}

		if err = json.NewEncoder(w).Encode(entity); err != nil {
			return err
		}

		return nil
	}
}

func main() {
	repository := &memoryRepository{entities: map[int64]Entity{
		1: {1, "One"},
		2: {2, "Two"},
		3: {3, "Three"},
	}}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /entities", HandleStatus(GetEntityByID(repository)))

	if err := http.ListenAndServe(":8080", mux); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
