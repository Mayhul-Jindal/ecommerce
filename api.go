package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal.git/types"
)

// type for my api handlers
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// helper function for writing responses
func writeJSON(w http.ResponseWriter, s int, resp types.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(resp)
}

// helper function for handling errors at a single place (increases maintainability)
func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		reqValue := r.RemoteAddr
		ctx = context.WithValue(ctx, types.CtxKey, reqValue)

		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, types.Response{Title: "error aa gayi pancho"})
		}
	}
}

type APIServer struct {
	listenAddr string
	svc        BookManager
}

func NewAPIServer(listenAddr string, svc BookManager) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *APIServer) Run() {
	http.HandleFunc("/", makeAPIFunc(s.handleRoot))

	if err := http.ListenAndServe(s.listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}

func (s *APIServer) handleRoot(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp := types.Response{
		Title: "this is a test",
	}
	s.svc.Search(ctx,"hello")
	return writeJSON(w, http.StatusOK, resp)
}
