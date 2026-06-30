package dto

type (
	// PasswordLoginRequest is submitted by the login form for username/password
	// authentication, forwarded to Auth0's Resource Owner Password Grant.
	PasswordLoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	// GoogleLoginCallbackRequest carries the parameters Auth0 redirects back
	// with after a Google federated login.
	GoogleLoginCallbackRequest struct {
		Code  string `query:"code" validate:"required"`
		State string `query:"state" validate:"omitempty"`
	}

	// Auth0TokenResponse is the response body from Auth0's /oauth/token endpoint.
	Auth0TokenResponse struct {
		AccessToken string `json:"access_token"`
		IdToken     string `json:"id_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	// Auth0UserInfo is the response body from Auth0's /userinfo endpoint.
	Auth0UserInfo struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	// AuthResult is returned by the auth service after a successful login.
	AuthResult struct {
		UserID uint64
		Name   string
		Email  string
		RoleID uint64
		Token  string
	}
)
