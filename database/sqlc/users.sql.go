// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "Users" (
  username, email, hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING id, username, email, hashed_password, password_changed_at, is_admin, is_active, deactivated_at, is_deleted, deleted_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.IsAdmin,
		&i.IsActive,
		&i.DeactivatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, username, email, hashed_password, password_changed_at, is_admin, is_active, deactivated_at, is_deleted, deleted_at, created_at from "Users"
where id = $1 limit 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.IsAdmin,
		&i.IsActive,
		&i.DeactivatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}
