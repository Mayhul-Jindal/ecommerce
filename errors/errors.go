package errors

import "errors"

var (
	ErrorMethodNotAllowed = errors.New("method not allowed")
	ErrorBadRequest = errors.New("only json is supported")
)