package gatewayservice

import (
	"context"
	"log"
	"net/http"
	"strings"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)

func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), types.RemoteAddress, r.RemoteAddr)
		ctx = context.WithValue(ctx, types.UserAgent, r.UserAgent())
	
		err := fn(ctx, w, r)
		if err == nil {
			return
		}
		switch err {
		case errs.ErrorMethodNotAllowed:
			writeJSON(w, http.StatusMethodNotAllowed, r.URL.String(), APIError{Error: err.Error()})

		case errs.ErrorBadRequest:
			writeJSON(w, http.StatusBadRequest, r.URL.String(), APIError{Error: err.Error()})

		case errs.ErrorInvalidToken, errs.ErrorExpiredToken, errs.ErrorUnauthorized, errs.ErrorNoAuthHeader, errs.ErrorInvalidAuthHeader, errs.ErrorUnsupportedAuthType, errs.ErrorNotAuthorized, errs.ErrorExpiredSession:
			writeJSON(w, http.StatusUnauthorized, r.URL.String(), APIError{Error: err.Error()})

		case errs.ErrorAmountMismatch, errs.ErrorValidationFailed, errs.ErrorPayementNotVerified, errs.ErrorFileLimitExceeded:
			writeJSON(w, http.StatusBadRequest, r.URL.String(), APIError{Error: err.Error()})

		case errs.ErrorPageNotFound, errs.ErrorRecordNotFound:
			writeJSON(w, http.StatusNotFound, r.URL.String(), APIError{Error: err.Error()})

		default:
			if errs.ErrorCode(err) == errs.UniqueViolation || errs.ErrorCode(err) == errs.ForeignKeyViolation {
				writeJSON(w, http.StatusForbidden, r.URL.String(), APIError{Error: errs.ErrorUniqueOrForeignKeyViolation.Error()})
				return
			}

			writeJSON(w, http.StatusInternalServerError, r.URL.String(), APIError{Error: err.Error()})
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
		payload, err := s.tokenSvc.VerifyToken(accessToken)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, types.AuthorizationPayload, payload)
		return fn(ctx, w, r)
	}
}
