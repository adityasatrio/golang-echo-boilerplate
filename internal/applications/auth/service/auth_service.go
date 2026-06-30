package service

import (
	"context"
	"myapp/internal/applications/auth/dto"
)

// Claims are the custom JWT claims issued for an authenticated app session.
type Claims struct {
	UserID uint64 `json:"user_id"`
	RoleID uint64 `json:"role_id"`
	Email  string `json:"email"`
}

type AuthService interface {
	// LoginWithPassword authenticates against Auth0's Resource Owner Password
	// Grant, then finds-or-creates the local user record and issues an app JWT.
	LoginWithPassword(ctx context.Context, request *dto.PasswordLoginRequest) (*dto.AuthResult, error)

	// BuildGoogleAuthorizeURL builds the Auth0 redirect URL to start a Google
	// federated login.
	BuildGoogleAuthorizeURL(state string) string

	// HandleGoogleCallback exchanges the authorization code returned by Auth0,
	// fetches the user's profile, then finds-or-creates the local user record
	// and issues an app JWT.
	HandleGoogleCallback(ctx context.Context, code string) (*dto.AuthResult, error)

	// IssueJWT signs a short-lived app session token for the given user.
	IssueJWT(userID uint64, roleID uint64, email string) (string, error)

	// VerifyJWT validates an app session token and returns its claims.
	VerifyJWT(tokenString string) (*Claims, error)
}
