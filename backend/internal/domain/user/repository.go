package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrEmailExists   = errors.New("email already exists")
	ErrUsernameTaken = errors.New("username taken")
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByGoogleID(ctx context.Context, googleID string) (*User, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
}
