package auth0

import (
	"context"
	"myapp/internal/applications/auth/dto"
)

// Auth0Client wraps calls to the Auth0 Authentication API.
type Auth0Client interface {
	// LoginWithPassword exchanges an email/password pair for an Auth0 token via
	// the Resource Owner Password Grant.
	LoginWithPassword(ctx context.Context, email string, password string) (*dto.Auth0TokenResponse, error)

	// BuildGoogleAuthorizeURL builds the Auth0 /authorize redirect URL for the
	// Google federated connection.
	BuildGoogleAuthorizeURL(state string) string

	// ExchangeCodeForToken exchanges an authorization code (from the Google
	// login callback) for an Auth0 token via the Authorization Code grant.
	ExchangeCodeForToken(ctx context.Context, code string) (*dto.Auth0TokenResponse, error)

	// GetUserInfo fetches the Auth0 profile (email, name, picture) for the
	// given access token.
	GetUserInfo(ctx context.Context, accessToken string) (*dto.Auth0UserInfo, error)
}
