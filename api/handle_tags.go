package api

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

// handle tags
func (s *APIServer) handleTags(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	var err error

	switch action {
	case "get":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		resp, err = s.bookSvc.GetAllTags(ctx)
		if err != nil {
			return err
		}

	case "add":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.CreateTagRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.CreateTag(ctx, req)
		if err != nil {
			return err
		}

	case "update":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.UpdateTagRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.bookSvc.UpdateTag(ctx, req)
		if err != nil {
			return err
		}

	case "delete":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.DeleteTagRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.bookSvc.DeleteTag(ctx, req)
		if err != nil {
			return err
		}

		resp = map[string]string{"message": "ok"}
	default:
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
