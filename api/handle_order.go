package api

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "place":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.PlaceOrderRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.PlaceOrder(ctx, req)
		if err != nil {
			return err
		}

	case "verify":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.VerifyOrderRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.VerifyOrder(ctx, req)
		if err != nil {
			return err
		}

	case "get":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.GetPurchasesRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.GetPurchases(ctx, req)
		if err != nil {
			return err
		}

	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
