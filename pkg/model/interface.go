package model

import (
	"context"

	"getitqec.com/server/user/pkg/dto"
)

// UserModelI APIs
type UserModelI interface {
	// Google
	GoogleSignIn(ctx context.Context, idToken string, accessToken string) (*dto.User, bool, error)

	// User
	GetUser(ctx context.Context, id string) (*dto.User, error)

	SearchUser(ctx context.Context, email string) (*dto.User, error)
}

// PortfolioModelI APIs
type PortfolioModelI interface {
}
