package auth

import (
	"context"
	"errors"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/password"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	usersRepo user.UserRepository
	jwt       *JWTManager
}

func NewAuthService(userRepo user.UserRepository, jwtManager *JWTManager) *AuthService {
	return &AuthService{
		usersRepo: userRepo,
		jwt:       jwtManager,
	}
}

func (s *AuthService) Login(ctx context.Context, loginInfo LoginInfo) (TokenPair, error) {
	u, err := s.usersRepo.GetUserByEmail(ctx, loginInfo.Email)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return TokenPair{}, ErrInvalidCredentials
		}
		return TokenPair{}, err
	}
	if err := password.ComparePassword(loginInfo.Password, u.HashedPassword); err != nil {
		return TokenPair{}, ErrInvalidCredentials
	}

	access, err := s.jwt.GenerateAccessToken(u.ID)
	if err != nil {
		return TokenPair{}, err
	}
	refresh, err := s.jwt.GenerateRefreshToken(u.ID)
	if err != nil {
		return TokenPair{}, err
	}
	return TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
