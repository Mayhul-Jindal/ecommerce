// types for auth microservice
package types

import (
	"time"

	"github.com/google/uuid"
)

type contextKey string

const (
	RemoteAddress        contextKey = "remote_address"
	UserAgent            contextKey = "user_agent"
	AuthorizationPayload contextKey = "authorization_payload"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=7"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7"`
}

type UserResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	ID       int64  `json:"id"`
	Username string `json:"username" validate:"required,min=7"`
	Password string `json:"password" validate:"required,min=7"`
}

type LoginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  UserResponse `json:"user"`
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type VerifyEmailResponse struct {
	IsVerified bool `json:"is_verified"`
}