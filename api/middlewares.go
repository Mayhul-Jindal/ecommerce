package api

import (
	"context"
	"log"
	"net/http"
	"strings"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)

// type for api handlers
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// type for api error
type APIError struct {
	Error string `json:"error"`
}

// centralized error handling
func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), types.RemoteAddress, r.RemoteAddr)
		ctx = context.WithValue(ctx, types.UserAgent, r.UserAgent())
		ctx = context.WithValue(ctx, types.Route, r.URL.String())
		ctx = context.WithValue(ctx, types.Method, r.Method)
		
		err := fn(ctx, w, r)
		if err == nil {
			return
		}

		switch err {
		case errs.ErrorMethodNotAllowed:
			writeJSON(ctx, w, http.StatusMethodNotAllowed, APIError{Error: err.Error()})

		case errs.ErrorBadRequest:
			writeJSON(ctx, w, http.StatusBadRequest, APIError{Error: err.Error()})

		case errs.ErrorInvalidToken, errs.ErrorExpiredToken, errs.ErrorUnauthorized, errs.ErrorNoAuthHeader, errs.ErrorInvalidAuthHeader, errs.ErrorUnsupportedAuthType, 
			errs.ErrorNotAuthorized, errs.ErrorExpiredSession, errs.ErrorEmailNotVerified:
			writeJSON(ctx, w, http.StatusUnauthorized, APIError{Error: err.Error()})

		case errs.ErrorAmountMismatch, errs.ErrorValidationFailed, errs.ErrorPayementNotVerified, errs.ErrorFileLimitExceeded:
			writeJSON(ctx, w, http.StatusBadRequest, APIError{Error: err.Error()})

		case errs.ErrorPageNotFound, errs.ErrorRecordNotFound, errs.ErrorEmailNotVerified:
			writeJSON(ctx, w, http.StatusNotFound, APIError{Error: err.Error()})

		default:
			if errs.ErrorCode(err) == errs.UniqueViolation || errs.ErrorCode(err) == errs.ForeignKeyViolation {
				writeJSON(ctx, w, http.StatusForbidden, APIError{Error: errs.ErrorUniqueOrForeignKeyViolation.Error()})
				return
			}

			writeJSON(ctx, w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

// middle ware for checking authorization
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
		payload, err := s.token.VerifyToken(accessToken)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, types.AuthorizationPayload, payload)
		return fn(ctx, w, r)
	}
}
