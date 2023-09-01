// TODO
/*
- I have is_email_verified  {usethis for add to cart restriction}
- I have is_admin
- I have is_deactivated
- T have is_deleted
*/

// this is currentyl acting as api gateway for my whole architecture
package gatewayservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/authService"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/bookService"
	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// this helps me to decouple the auth logic from book management stuff
type APIServer struct {
	listenAddr string
	bookSvc    bookService.Manager
	authSvc    authService.Manager
	tokenSvc   token.Maker
	validator  *validator.Validate
}

// type for api handlers
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// type for api error
type APIError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, authSvc authService.Manager, bookSvc bookService.Manager, tokenSvc token.Maker, validator *validator.Validate) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		bookSvc:    bookSvc,
		authSvc:    authSvc,
		tokenSvc:   tokenSvc,
		validator:  validator,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// these are some of the global erros
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, r.URL.String(), APIError{Error: errs.ErrorPageNotFound.Error()})
	})


	router.HandleFunc("/users/{action}", makeAPIFunc(s.handleUsers))
	router.HandleFunc("/users/verify_email", makeAPIFunc(s.handleVerifyEmail))
	router.HandleFunc("/search", makeAPIFunc(s.handleSearch))
	router.HandleFunc("/tags", makeAPIFunc(s.handleTags))
	router.HandleFunc("/recommendations/hotselling", makeAPIFunc(s.handleHotSelling))

	router.HandleFunc("/book/{action}", makeAPIFunc(s.checkAuthorization(s.handleBook)))
	router.HandleFunc("/cart/{action}", makeAPIFunc(s.checkAuthorization(s.handleCart)))
	router.HandleFunc("/order/{action}", makeAPIFunc(s.checkAuthorization(s.handleOrder)))
	router.HandleFunc("/review/{action}", makeAPIFunc(s.checkAuthorization(s.handleReviews)))
	router.HandleFunc("/recommendations/personal", makeAPIFunc(s.checkAuthorization(s.handleRecommendations)))

	// todo Add a transaction for deletion of record
	// router.HandleFunc("/users/deactivate", makeAPIFunc(s.checkAuthorization(s.handleDeactivateAccount)))
	// router.HandleFunc("/users/delete", makeAPIFunc(s.checkAuthorization(s.handleDeleteAccount)))

	// timeouts are set for the request
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1%s", s.listenAddr),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// helper function for writing responses
// this also logs the responses
func writeJSON(w http.ResponseWriter, s int, route string, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	log.Printf("status:%v route:%v", s, route)
	return json.NewEncoder(w).Encode(v)
}
