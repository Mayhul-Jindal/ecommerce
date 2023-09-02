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
	"net/http"
	"os"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/authService"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/bookService"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
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

var logger zerolog.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Logger()

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

	router.HandleFunc("/users/{action}", makeAPIFunc(s.handleUsers))
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

	router.HandleFunc("/{random}", makeAPIFunc(s.handlePageNotFound))

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1%s", s.listenAddr),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

// centralized logging for json api gateway
func writeJSON(ctx context.Context, w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	if apiErr, ok := v.(APIError); ok {
		logger.Error().
			Int("status", s).
			Str("route", ctx.Value(types.Route).(string)).
			Str("method", ctx.Value(types.Method).(string)).
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Str("err", apiErr.Error).
			Send()
			
	} else {
		logger.Info().
			Int("status", s).
			Str("route", ctx.Value(types.Route).(string)).
			Str("method", ctx.Value(types.Method).(string)).
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Send()
	}

	return json.NewEncoder(w).Encode(v)
}
