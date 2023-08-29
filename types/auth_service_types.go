// types for auth microservice
package types

import "time"

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
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}