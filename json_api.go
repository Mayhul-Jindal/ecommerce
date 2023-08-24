// TODO
/*
- make different endpoints for different things
- make tests for http endpoints
*/

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

type Response struct{
	Title string `json:"title"`
}

type APIServer struct {
	listenAddr string
	db         Storer
}

func NewAPIServer(listenAddr string, db Storer) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", makeAPIFunc(s.handleRoot))
	router.HandleFunc("/users/signup", makeAPIFunc(s.handleSignup)).Methods("POST")
	router.HandleFunc("/users/login", makeAPIFunc(s.handleLogin)).Methods("POST")
	router.HandleFunc("/users", makeAPIFunc(s.handleUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", makeAPIFunc(s.handleUserById)).Methods("GET")

	http.ListenAndServe(s.listenAddr, router)
}

// handle for root
func (s *APIServer) handleRoot(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return writeJSON(w, http.StatusMethodNotAllowed, APIError{Error: errors.ErrorMethodNotAllowed.Error()})
	}

	resp := Response{Title: "hello wolrd"}
	return writeJSON(w, http.StatusOK, resp)
}

// handle for signup
func (s *APIServer) handleSignup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp := Response{Title: "hello wolrd"}
	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleLogin(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp := Response{Title: "hello wolrd"}
	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp := Response{Title: "hello wolrd"}
	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUserById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idd, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("1")
		return err
	}

	resp, err := s.db.GetBook(ctx, idd)
	if err != nil {
		log.Println("2")
		return err
	}

	log.Printf("%+v", resp)
	return writeJSON(w, http.StatusOK, resp)
}












// type for my api handlers
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// helper function for writing responses
func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}

type APIError struct {
	Error string `json:"error"`
}

// helper function for handling errors at a single place (increases maintainability)
func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	// this context help us the manage the life-cycle of a request
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		// max 3 seconds for a request
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		ctx = context.WithValue(ctx, types.CtxKey, r.RemoteAddr)
		defer cancel()

		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}
