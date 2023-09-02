package gatewayservice

import (
	"context"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
)

func (s *APIServer) handlePageNotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return errs.ErrorPageNotFound
}
