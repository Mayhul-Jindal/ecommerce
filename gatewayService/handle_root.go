package gatewayservice

import (
	"context"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
)

func (s *APIServer) handleRoot(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	resp := map[string]string{"message": "ok"}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
