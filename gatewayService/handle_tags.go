package gatewayservice

import (
	"context"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
)

// handle tags
func (s *APIServer) handleTags(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	// request validation
	resp, err := s.bookSvc.GetAllTags(ctx)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, r.URL.String(), resp)
}
