package api

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

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

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
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

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
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

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.bookSvc.DeleteCartItem(ctx, req)
		if err != nil {
			return err
		}

	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
