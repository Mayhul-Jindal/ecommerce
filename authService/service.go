package authService

import (
	"context"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/worker"
	"github.com/jackc/pgx/v5/pgtype"
)

type Manager interface {
	// unauthorized
	SignUp(ctx context.Context, req types.CreateUserRequest) (types.UserResponse, error)
	Login(ctx context.Context, req types.LoginUserRequest) (types.LoginUserResponse, error)
	RenewAccess(ctx context.Context, req types.RenewAccessTokenRequest) (types.RenewAccessTokenResponse, error)
	VerifyEmail(ctx context.Context, id int64, secret_code string) (types.VerifyEmailResponse, error) 
	ResendEmail(ctx context.Context, req types.ResendEmailRequest)

	// authorized
	DeactivateAccount(ctx context.Context, req types.DeactivateAccountRequest) (types.DeactivateAccountResponse, error)
	DeleteAccount(ctx context.Context, req types.DeleteAccountRequest) (types.DeleteAccountResponse, error)
}

type authManager struct {
	config     util.Config
	tokenMaker token.Maker
	db         database.Storer
	worker     worker.Worker
}

func NewManager(config util.Config, tokenMaker token.Maker, db database.Storer, worker worker.Worker) Manager {
	return &authManager{
		config:     config,
		tokenMaker: tokenMaker,
		db:         db,
		worker:     worker,
	}
}

func (a *authManager) SignUp(ctx context.Context, req types.CreateUserRequest) (types.UserResponse, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return types.UserResponse{}, err
	}

	user, err := a.db.GetUserByUsername(ctx, req.Username)
	if err == nil {
		if  time.Now().Before(user.DeletedAt.AddDate(0,0,15)) {
			return types.UserResponse{}, errs.ErrorAccountIsDeleted
		}
	}
	
	createUserParams := database.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	arg := database.CreateUserTxParams{
		CreateUserParams: createUserParams,
		AfterCreate:      a.worker.EnqueueSendVerifyEmail,
	}

	resp, err := a.db.CreateUserTx(ctx, arg)
	if err != nil {
		return types.UserResponse{}, err
	}

	return newUserResponse(resp.User), nil
}

func (a *authManager) Login(ctx context.Context, req types.LoginUserRequest) (types.LoginUserResponse, error) {
	params := database.GetUserParams{
		ID:       req.UserID,
		Username: req.Username,
	}

	user, err := a.db.GetUser(ctx, params)
	if err != nil {
		return types.LoginUserResponse{}, errs.ErrorNoUser
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return types.LoginUserResponse{}, errs.ErrorUnauthorized
	}

	if !user.IsActive {
		return types.LoginUserResponse{}, errs.ErrorAccountIsDeactivated
	}


	accessToken, accessPayload, err := a.tokenMaker.CreateToken(
		user.ID,
		a.config.ACCESS_TOKEN_DURATION,
	)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	refreshToken, refreshPayload, err := a.tokenMaker.CreateToken(
		user.ID,
		a.config.REFRESH_TOKEN_DURATION,
	)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	sessionParams := database.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Value(types.UserAgent).(string),
		ClientIp:     ctx.Value(types.RemoteAddress).(string),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}
	session, err := a.db.CreateSession(ctx, sessionParams)
	if err != nil {
		return types.LoginUserResponse{}, err
	}

	resp := types.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	return resp, nil
}

func (a *authManager) RenewAccess(ctx context.Context, req types.RenewAccessTokenRequest) (types.RenewAccessTokenResponse, error) {
	refreshPayload, err := a.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return types.RenewAccessTokenResponse{}, err
	}
	
	session, err := a.db.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return types.RenewAccessTokenResponse{}, err
	}

	if session.IsBlocked {
		return types.RenewAccessTokenResponse{}, errs.ErrorUnauthorized
	}

	if session.UserID != refreshPayload.UserID {
		return types.RenewAccessTokenResponse{}, errs.ErrorUnauthorized
	}

	if session.RefreshToken != req.RefreshToken {
		return types.RenewAccessTokenResponse{}, errs.ErrorUnauthorized
	}

	if time.Now().After(session.ExpiresAt) {
		return types.RenewAccessTokenResponse{}, errs.ErrorExpiredSession
	}

	accessToken, accessPayload, err := a.tokenMaker.CreateToken(
		refreshPayload.UserID,
		a.config.ACCESS_TOKEN_DURATION,
	)
	if err != nil {
		return types.RenewAccessTokenResponse{}, err
	}

	resp := types.RenewAccessTokenResponse{
		UserID: refreshPayload.UserID,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return resp, nil
}

func (a *authManager) VerifyEmail(ctx context.Context, id int64, secret_code string) (types.VerifyEmailResponse, error) {
	txResult, err := a.db.VerifyEmailTx(ctx, database.VerifyEmailTxParams{
		ID:         id,
		SecretCode: secret_code,
	})
	if err != nil {
		return types.VerifyEmailResponse{}, err
	}

	resp := types.VerifyEmailResponse{
		UserID: txResult.VerifyEmail.UserID,
		IsVerified: txResult.User.IsEmailVerified,
	}
	
	return resp, nil
}

func (a *authManager) ResendEmail(ctx context.Context, req types.ResendEmailRequest) {
	a.worker.EnqueueSendVerifyEmail(database.GetUserParams{
		ID: req.UserID,
		Username: req.Username,
	})
}


func (a *authManager) DeactivateAccount(ctx context.Context, req types.DeactivateAccountRequest) (types.DeactivateAccountResponse, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return types.DeactivateAccountResponse{}, errs.ErrorNotAuthorized
	}

	user, err := a.db.GetUserById(ctx, authPayload.UserID)
	if err != nil {
		return types.DeactivateAccountResponse{}, err
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return types.DeactivateAccountResponse{}, errs.ErrorUnauthorized
	}

	_, err = a.db.UpdateUser(ctx, database.UpdateUserParams{
		ID: authPayload.UserID,
		IsActive: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
		DeactivatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return types.DeactivateAccountResponse{}, err
	}

	return types.DeactivateAccountResponse{
		Message: "Account Deactivated Successfully",
	}, nil
}

// core deletion logic should move to a worker
func (a *authManager) DeleteAccount(ctx context.Context, req types.DeleteAccountRequest) (types.DeleteAccountResponse, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return types.DeleteAccountResponse{}, errs.ErrorNotAuthorized
	}

	user, err := a.db.GetUserById(ctx, authPayload.UserID)
	if err != nil {
		return types.DeleteAccountResponse{}, err
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return types.DeleteAccountResponse{}, errs.ErrorUnauthorized
	}

	_, err = a.db.UpdateUser(ctx, database.UpdateUserParams{
		ID: authPayload.UserID,
		IsDeleted: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		return types.DeleteAccountResponse{}, err
	}

	a.worker.EnqueueDeleteOperation(authPayload.UserID)

	return types.DeleteAccountResponse{
		Message: "Account Deleted Successfully",
	}, nil
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


