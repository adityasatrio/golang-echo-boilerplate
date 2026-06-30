package auth0

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
	"myapp/configs/credential"
	"myapp/configs/http"
	"myapp/internal/applications/auth/dto"
)

type Auth0ClientImpl struct {
	ClientApi *resty.Client
}

func NewAuth0Client() *Auth0ClientImpl {
	return &Auth0ClientImpl{
		ClientApi: http.New(),
	}
}

func (a *Auth0ClientImpl) domainURL() string {
	return fmt.Sprintf("https://%s", credential.GetString("auth0.domain"))
}

func (a *Auth0ClientImpl) LoginWithPassword(ctx context.Context, email string, password string) (*dto.Auth0TokenResponse, error) {
	body := map[string]string{
		"grant_type":    "password",
		"username":      email,
		"password":      password,
		"client_id":     credential.GetString("auth0.clientId"),
		"client_secret": credential.GetString("auth0.clientSecret"),
		"audience":      credential.GetString("auth0.audience"),
		"scope":         "openid profile email",
	}

	response := &dto.Auth0TokenResponse{}
	resp, err := a.ClientApi.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(response).
		Post(a.domainURL() + "/oauth/token")

	if err != nil {
		log.Errorf("auth0 password login http error: %v", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		errMsg := fmt.Sprintf("auth0 password login failed with status %d: %s", resp.StatusCode(), resp.String())
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	return response, nil
}

func (a *Auth0ClientImpl) BuildGoogleAuthorizeURL(state string) string {
	values := url.Values{}
	values.Set("response_type", "code")
	values.Set("client_id", credential.GetString("auth0.clientId"))
	values.Set("redirect_uri", credential.GetString("auth0.callbackUrl"))
	values.Set("connection", "google-oauth2")
	values.Set("scope", "openid profile email")
	values.Set("state", state)

	return fmt.Sprintf("%s/authorize?%s", a.domainURL(), values.Encode())
}

func (a *Auth0ClientImpl) ExchangeCodeForToken(ctx context.Context, code string) (*dto.Auth0TokenResponse, error) {
	body := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     credential.GetString("auth0.clientId"),
		"client_secret": credential.GetString("auth0.clientSecret"),
		"code":          code,
		"redirect_uri":  credential.GetString("auth0.callbackUrl"),
	}

	response := &dto.Auth0TokenResponse{}
	resp, err := a.ClientApi.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(response).
		Post(a.domainURL() + "/oauth/token")

	if err != nil {
		log.Errorf("auth0 code exchange http error: %v", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		errMsg := fmt.Sprintf("auth0 code exchange failed with status %d: %s", resp.StatusCode(), resp.String())
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	return response, nil
}

func (a *Auth0ClientImpl) GetUserInfo(ctx context.Context, accessToken string) (*dto.Auth0UserInfo, error) {
	response := &dto.Auth0UserInfo{}
	resp, err := a.ClientApi.R().
		SetContext(ctx).
		SetAuthToken(accessToken).
		SetResult(response).
		Get(a.domainURL() + "/userinfo")

	if err != nil {
		log.Errorf("auth0 userinfo http error: %v", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		errMsg := fmt.Sprintf("auth0 userinfo failed with status %d: %s", resp.StatusCode(), resp.String())
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	return response, nil
}
