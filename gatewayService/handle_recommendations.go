package gatewayservice

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)

func (s *APIServer) handleHotSelling(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errs.ErrorMethodNotAllowed
	}

	var req types.GetHotSellingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(1)
		return errs.ErrorBadRequest
	}

	err = s.validator.Struct(req)
	if err != nil {
		log.Println(1)
		return errs.ErrorBadRequest
	}

	resp, err := s.bookSvc.GetHotSelling(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}

func (s *APIServer) handleRecommendations(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var req types.GetPersonalRecommendationsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	err = s.validator.Struct(req)
	if err != nil {
		return err
	}

	resp, err := s.bookSvc.GetPersonalRecommendations(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}
