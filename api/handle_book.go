package api

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleBook(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "get":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.GetBookRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.GetBook(ctx, req)
		if err != nil {
			return err
		}

	// admin hee access kar payega yeh
	case "add":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.AddBookRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.AddBook(ctx, req)
		if err != nil {
			return err
		}

	case "update":
		if r.Method != "PATCH" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.UpdateBookRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.UpdateBook(ctx, req)
		if err != nil {
			return err
		}

	// case "delete":
	// 	if r.Method != "DELETE" {
	// 		return errs.ErrorMethodNotAllowed
	// 	}

	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
