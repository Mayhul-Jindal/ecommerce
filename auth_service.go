// TODO
/*
- this can be moved to a new microservice
- give error types help us to identify what to send at what error if that makes sense
*/

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
)

type AuthManager interface {
	SignUp(context.Context, *http.Request) (UserResponse, error)
}

type authManager struct {
	db Storer
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required|min_len:7"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required|min_len:7"`
}

type UserResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user database.User) UserResponse {
	return UserResponse{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func NewAuthManager(db Storer) AuthManager {
	return &authManager{
		db: db,
	}
}

// signup logic here
func (a *authManager) SignUp(ctx context.Context, r *http.Request) (UserResponse, error) {
	// request validation goes here
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return UserResponse{}, err
	}

	// here the logic with database and stuff is done here
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return UserResponse{}, err
	}

	user := database.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	registeredUser, err := a.db.CreateUser(ctx, user)
	if err != nil {
		return UserResponse{}, err
	}

	return newUserResponse(registeredUser), nil
}
