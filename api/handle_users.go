package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	action := mux.Vars(r)["action"]

	var resp any
	switch action {
	case "signup":
		if r.Method != "POST" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.authSvc.SignUp(ctx, req)
		if err != nil {
			return err
		}

	case "login":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.LoginUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.authSvc.Login(ctx, req)
		if err != nil {
			return err
		}

	case "renew_access":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.RenewAccessTokenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		resp, err = s.authSvc.RenewAccess(ctx, req)
		if err != nil {
			return err
		}

	case "resend_email":
		if r.Method != "GET" {
			return errs.ErrorMethodNotAllowed
		}

		var req types.ResendEmailRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		err = s.validator.Struct(req)
		if err != nil {
			return errs.ErrorBadRequest
		}

		s.authSvc.ResendEmail(ctx, req)

		resp = map[string]string{"message": "resending email"}

	case "verify_email":
		num, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

		if err != nil || num < 1 {
			fmt.Println("Error:", err)
			return errs.ErrorBadRequest
		}
		secretCode := r.URL.Query().Get("secret_code")

		resp, err = s.authSvc.VerifyEmail(ctx, num, secretCode)
		if err != nil {
			return err
		}

	default:
		return errs.ErrorPageNotFound
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}


func (s *APIServer) handleDeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "DELETE" {
		return errs.ErrorMethodNotAllowed
	}

	var req types.DeleteAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	err = s.validator.Struct(req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	resp, err := s.authSvc.DeleteAccount(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}


func (s *APIServer) handleDeactivateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "PATCH" {
		return errs.ErrorMethodNotAllowed
	}

	var req types.DeactivateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	err = s.validator.Struct(req)
	if err != nil {
		return errs.ErrorBadRequest
	}

	resp, err := s.authSvc.DeactivateAccount(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(ctx, w, http.StatusOK, resp)
}