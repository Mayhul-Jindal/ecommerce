package api

import (
	"context"
	"encoding/json"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)


func (s *APIServer) handleSearch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	// request validation
	var req types.SearchBooksV1Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	err = s.validator.Struct(req)
	if err != nil {
		return err
	}

	resp, err := s.bookSvc.Search(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
