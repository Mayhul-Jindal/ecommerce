// TODO
/*
- make different endpoints for different things
- make tests for http endpoints
*/

// this is currentyl acting as api gateway for my whole architecture
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// test response
type Response struct {
	Title string `json:"title"`
}

// this helps me to decouple the auth logic from book management stuff
type APIServer struct {
	listenAddr string
	bookSvc    BookManager
	authSvc    AuthManager
	tokenSvc   token.Maker
}

func NewAPIServer(authSvc AuthManager, bookSvc BookManager, tokenSvc token.Maker) *APIServer {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	return &APIServer{
		listenAddr: config.SERVER_PORT,
		bookSvc:    bookSvc,
		authSvc:    authSvc,
		tokenSvc:   tokenSvc,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// unauthorized routes
	router.HandleFunc("/users/signup", makeAPIFunc(s.handleSignup))
	router.HandleFunc("/users/login", makeAPIFunc(s.handleLogin))
	router.HandleFunc("/search", makeAPIFunc(s.handleSearch))

	// authorized routes
	router.HandleFunc("/", makeAPIFunc(s.checkAuthorization(s.handleRoot)))
	// router.HandleFunc("/users/deactivate", makeAPIFunc(s.checkAuthorization(s.handle)))
	// router.HandleFunc("/users/delete", makeAPIFunc(s.checkAuthorization(s.handleRoot)))
	router.HandleFunc("/cart/{action}", makeAPIFunc(s.checkAuthorization(s.handleCart)))
	router.HandleFunc("/order/{action}", makeAPIFunc(s.checkAuthorization(s.handleOrder)))
	// router.HandleFunc("/review", makeAPIFunc(s.checkAuthorization(s.handleCart)))

	http.ListenAndServe(s.listenAddr, router)
}

// handle signup
func (s *APIServer) handleSignup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errs.ErrorMethodNotAllowed
	}

	// request validation
	var req types.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	resp, err := s.authSvc.SignUp(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, resp)
}

// handle login
func (s *APIServer) handleLogin(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errs.ErrorMethodNotAllowed
	}

	// request validation
	var req types.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	resp, err := s.authSvc.Login(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, resp)
}

// handle search
func (s *APIServer) handleSearch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	// request validation
	var req types.SearchBooksV1Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	resp, err := s.bookSvc.Search(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, resp)
}

// ----------------------------------- authourized routes from here -----------------------------------

// test for authorization
func (s *APIServer) handleRoot(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	resp := Response{
		Title: "hello world",
	}

	return writeJSON(w, http.StatusOK, resp)
}

// handle cart dynamically :/
func (s *APIServer) handleCart(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "get":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.GetCartRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return err
		}

		resp, err = s.bookSvc.GetCart(ctx, req)
		if err != nil {
			return err
		}

	case "add":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.AddToCartRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return err
		}

		resp, err = s.bookSvc.AddToCart(ctx, req)
		if err != nil {
			return err
		}

	case "delete":
		if r.Method != "DELETE" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.DeleteCartItemRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return err
		}

		err = s.bookSvc.DeleteCartItem(ctx, req)
		if err != nil {
			return err
		}

	default:
		return errs.ErrorBadRequest
	}

	return writeJSON(w, http.StatusOK, resp)
}

// handle order dynamically :/
func (s *APIServer) handleOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "place":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.PlaceOrderRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return err
		}

		resp, err = s.bookSvc.PlaceOrder(ctx, req)
		if err != nil {
			return err
		}

	default:
		return errs.ErrorBadRequest
	}

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

// helper function for handling errors at a single place (increases maintainability). I can also log here successfully(beta)
func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// max 3 seconds for a request
		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
		ctx = context.WithValue(ctx, types.RemoteAddress, r.RemoteAddr)
		defer cancel()
		r = r.WithContext(ctx)

		if err := fn(ctx, w, r); err != nil {

			
			// abhi currently mostly internal server hee dikhenge
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

// reason to keep it as a method of struct so that it has neat access token maker
func (s *APIServer) checkAuthorization(fn APIFunc) APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		authorizationHeader := r.Header.Get("authorization")

		if len(authorizationHeader) == 0 {
			return errs.ErrorNoAuthHeader
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.Println("3")
			return errs.ErrorInvalidAuthHeader
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			log.Println("4")
			return errs.ErrorUnsupportedAuthType
		}

		accessToken := fields[1]
		payload, err := s.tokenSvc.VerifyToken(accessToken)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, types.AuthorizationPayload, payload)
		return fn(ctx, w, r)
	}
}
