// TODO
/*
- this can be moved to a new microservice
- give error types help us to identify what to send at what error if that makes sense
- handle database errors in the database package itself
*/

package main

import (
	"context"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
)

type AuthManager interface {
	SignUp(ctx context.Context, req types.CreateUserRequest) (types.UserResponse, error)
	Login(ctx context.Context, req types.LoginUserRequest) (types.LoginUserResponse, error)
}

type authManager struct {
	config     util.Config
	tokenMaker token.Maker
	db         Storer
}

func NewAuthManager(config util.Config, tokenMaker token.Maker, db Storer) AuthManager {
	return &authManager{
		config:     config,
		tokenMaker: tokenMaker,
		db:         db,
	}
}

func (a *authManager) SignUp(ctx context.Context, req types.CreateUserRequest) (types.UserResponse, error) {
	// request validation goes here

	// here the logic with database and stuff is done here
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return types.UserResponse{}, err
	}

	params := database.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	registeredUser, err := a.db.CreateUser(ctx, params)
	if err != nil {
		return types.UserResponse{}, err
	}

	return newUserResponse(registeredUser), nil
}

func (a *authManager) Login(ctx context.Context, req types.LoginUserRequest) (types.LoginUserResponse, error) {
	params := database.GetUserParams{
		ID:       req.ID,
		Username: req.Username,
	}

	user, err := a.db.GetUser(ctx, params)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return types.LoginUserResponse{}, errs.ErrorUnauthorized
	}

	accessToken, _, err := a.tokenMaker.CreateToken(
		req.ID,
		a.config.ACCESS_TOKEN_DURATION,
	)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	resp := types.LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	return resp, nil
}

// helper functions
func newUserResponse(user database.User) types.UserResponse {
	return types.UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
