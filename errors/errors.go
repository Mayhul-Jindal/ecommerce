package errs

import "errors"

var (
	ErrorMethodNotAllowed    = errors.New("method not allowed")
	ErrorBadRequest          = errors.New("bad request")
	ErrorInvalidToken        = errors.New("token is invalid")
	ErrorExpiredToken        = errors.New("token has expired")
	ErrorUnauthorized        = errors.New("not authorized")
	ErrorNoAuthHeader        = errors.New("authorization header is not provided")
	ErrorInvalidAuthHeader   = errors.New("invalid authorization header format")
	ErrorUnsupportedAuthType = errors.New("unsupported authorization type")
	ErrorNotAuthorized       = errors.New("not authorized")
	ErrorAmountMismatch      = errors.New("amount mismatch retry again")
	ErrorValidationFailed    = errors.New("request validation is failed")
	ErrorExpiredSession      = errors.New("session expired")
)
