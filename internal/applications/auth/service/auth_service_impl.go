package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"myapp/configs/credential"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/auth/dto"
	"myapp/internal/applications/auth/repository/auth0"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	userRepository "myapp/internal/applications/user/repository"
	"myapp/internal/component/transaction"
)

// defaultAutoProvisionRoleID is assigned to users who first authenticate via
// Auth0 (password or Google) and have no existing local account. It matches
// the "User" role seeded by migrations/20260630120005_seed_roles_and_admin_user.go.
const defaultAutoProvisionRoleID = uint64(2)

type AuthServiceImpl struct {
	userRepository     userRepository.UserRepository
	roleUserRepository roleUserRepository.RoleUserRepository
	transaction        transaction.Trx
	auth0Client        auth0.Auth0Client
}

func NewAuthService(userRepository userRepository.UserRepository, roleUserRepository roleUserRepository.RoleUserRepository, transaction transaction.Trx, auth0Client auth0.Auth0Client) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository:     userRepository,
		roleUserRepository: roleUserRepository,
		transaction:        transaction,
		auth0Client:        auth0Client,
	}
}

func (s *AuthServiceImpl) LoginWithPassword(ctx context.Context, request *dto.PasswordLoginRequest) (*dto.AuthResult, error) {
	if _, err := s.auth0Client.LoginWithPassword(ctx, request.Email, request.Password); err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	user, err := s.findOrCreateUser(ctx, request.Email, request.Email)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	return s.buildAuthResult(user)
}

func (s *AuthServiceImpl) BuildGoogleAuthorizeURL(state string) string {
	return s.auth0Client.BuildGoogleAuthorizeURL(state)
}

func (s *AuthServiceImpl) HandleGoogleCallback(ctx context.Context, code string) (*dto.AuthResult, error) {
	token, err := s.auth0Client.ExchangeCodeForToken(ctx, code)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	profile, err := s.auth0Client.GetUserInfo(ctx, token.AccessToken)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	if profile.Email == "" {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, errors.New("google profile did not return an email"))
	}

	user, err := s.findOrCreateUser(ctx, profile.Email, profile.Name)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	return s.buildAuthResult(user)
}

// findOrCreateUser looks up a local user by email, auto-provisioning one with
// the default "User" role on first login. Auth0 is the source of truth for
// credentials, so the local password is a random value that is never used to
// authenticate directly.
func (s *AuthServiceImpl) findOrCreateUser(ctx context.Context, email string, name string) (*ent.User, error) {
	existing, err := s.userRepository.GetByEmail(ctx, email)
	if err == nil && existing != nil {
		return existing, nil
	}

	randomPassword, err := generateRandomPassword()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if name == "" {
		name = email
	}

	var created *ent.User
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {
		newUser := ent.User{
			Name:       name,
			Email:      email,
			Password:   string(hashedPassword),
			Avatar:     "",
			RoleID:     defaultAutoProvisionRoleID,
			IsVerified: true,
		}

		userResult, err := s.userRepository.CreateTx(ctx, tx.Client(), newUser)
		if err != nil {
			return err
		}
		created = userResult

		roleUserRequest := ent.RoleUser{
			UserID: userResult.ID,
			RoleID: defaultAutoProvisionRoleID,
		}
		if _, err := s.roleUserRepository.CreateTx(ctx, tx.Client(), roleUserRequest); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return created, nil
}

func (s *AuthServiceImpl) buildAuthResult(user *ent.User) (*dto.AuthResult, error) {
	token, err := s.IssueJWT(user.ID, user.RoleID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResult{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		Token:  token,
	}, nil
}

func (s *AuthServiceImpl) IssueJWT(userID uint64, roleID uint64, email string) (string, error) {
	expiryMinutes := credential.GetInt("jwt.expiryMinutes")
	if expiryMinutes <= 0 {
		expiryMinutes = 60
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"role_id": roleID,
		"email":   email,
		"exp":     time.Now().Add(time.Duration(expiryMinutes) * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(credential.GetString("jwt.secret")))
}

func (s *AuthServiceImpl) VerifyJWT(tokenString string) (*Claims, error) {
	parsed, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(credential.GetString("jwt.secret")), nil
	})
	if err != nil || !parsed.Valid {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, err)
	}

	mapClaims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, exceptions.NewBusinessLogicError(exceptions.AuthInvalidCredentials, errors.New("invalid token claims"))
	}

	userID, _ := mapClaims["user_id"].(float64)
	roleID, _ := mapClaims["role_id"].(float64)
	email, _ := mapClaims["email"].(string)

	return &Claims{
		UserID: uint64(userID),
		RoleID: uint64(roleID),
		Email:  email,
	}, nil
}

func generateRandomPassword() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
