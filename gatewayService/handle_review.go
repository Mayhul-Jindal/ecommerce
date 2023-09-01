package gatewayservice

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

// handle reviews
func (s *APIServer) handleReviews(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "get":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.GetReviewsRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.GetReviews(ctx, req)
		if err != nil {
			return err
		}

	case "add":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.AddReviewRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.AddReview(ctx, req)
		if err != nil {
			return err
		}

	case "update":
		if r.Method != "PATCH" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.UpdateReviewRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.UpdateReview(ctx, req)
		if err != nil {
			return err
		}

	case "delete":
		if r.Method != "DELETE" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.DeleteReviewRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.bookSvc.DeleteReview(ctx, req)
		if err != nil {
			return err
		}

		resp = map[string]string{"message": "ok"}
	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(w, http.StatusOK, r.URL.String(), resp)
}
