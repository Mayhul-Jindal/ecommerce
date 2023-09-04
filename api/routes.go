// TODO
/*
- I have is_email_verified  {usethis for add to cart restriction}
- I have is_admin
- I have is_deactivated
- T have is_deleted
*/

// this is currentyl acting as api gateway for my whole architecture
package api

import (
	"context"
	"encoding/json"
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
	token      token.Maker
	validator  *validator.Validate
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
		token:      tokenSvc,
		validator:  validator,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", makeAPIFunc(s.handleRoot))

	router.HandleFunc("/users/delete", makeAPIFunc(s.checkAuthorization(s.handleDeleteUser)))
	router.HandleFunc("/users/deactivate", makeAPIFunc(s.checkAuthorization(s.handleDeactivateUser)))
	router.HandleFunc("/users/{action}", makeAPIFunc(s.handleUsers))
	router.HandleFunc("/search", makeAPIFunc(s.handleSearch))
	router.HandleFunc("/tags/{action}", makeAPIFunc(s.handleTags))
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

	logger.Info().Msgf("starting server on %s", s.listenAddr)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		logger.Fatal().Msgf("failed to start a server at port %s", s.listenAddr)
		os.Exit(1)
	}
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
