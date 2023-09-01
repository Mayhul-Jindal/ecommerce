package gatewayservice

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

// handle order dynamically :/
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
	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(w, http.StatusOK, r.URL.String(), resp)
}
