package errs

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var (
	ErrorMethodNotAllowed            = errors.New("method not allowed")
	ErrorBadRequest                  = errors.New("bad request")
	ErrorInvalidToken                = errors.New("token is invalid")
	ErrorExpiredToken                = errors.New("token has expired")
	ErrorUnauthorized                = errors.New("not authorized")
	ErrorNoAuthHeader                = errors.New("authorization header is not provided")
	ErrorInvalidAuthHeader           = errors.New("invalid authorization header format")
	ErrorUnsupportedAuthType         = errors.New("unsupported authorization type")
	ErrorNotAuthorized               = errors.New("not authorized")
	ErrorAmountMismatch              = errors.New("amount mismatch retry again")
	ErrorValidationFailed            = errors.New("request validation is failed")
	ErrorExpiredSession              = errors.New("session expired")
	ErrorPayementNotVerified         = errors.New("payement is not verified")
	ErrorFileLimitExceeded           = errors.New("file size cannot exceed 1MB")
	ErrorPageNotFound                = errors.New("404 page not found")
	ErrorRecordNotFound              = pgx.ErrNoRows
	ErrorUniqueOrForeignKeyViolation = errors.New("forbidden")
	ErrorRequestTimeout              = errors.New("request timeout")
	ErrorEmailNotVerified            = errors.New("email is not verified")
	ErrorLinkExpired                 = errors.New("gmail link is expired")
	ErrorBookAlreadyBought           = errors.New("book is laready bought")
	ErrorAccountIsDeactivated        = errors.New("account is deactivated for now try after 24 hrs to reactivate the account")
	ErrorAccountIsDeleted            = errors.New("you cannot use these credentials for 15 days")
	ErrorNoUser                      = errors.New("no such user found")
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

// refer to backend master class

// status forbidden
