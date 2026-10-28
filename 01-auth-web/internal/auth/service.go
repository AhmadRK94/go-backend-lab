package auth

import (
	"context"
	"errors"
	"strconv"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/password"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Service struct {
	usersRepo user.UserRepository
	jwt       *JWTManager
}

func NewService(userRepo user.UserRepository, jwtManager *JWTManager) *Service {
	return &Service{
		usersRepo: userRepo,
		jwt:       jwtManager,
	}
}

func (s *Service) Login(ctx context.Context, loginInfo LoginInfo) (TokenPair, error) {
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

func (s *Service) Refresh(refreshToken string) (string, error) {
	claims, err := s.jwt.ValidateToken(refreshToken, "refresh")
	if err != nil {
		return "", err
	}
	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil || id == 0 {
		return "", ErrInvalidToken
	}

	return s.jwt.GenerateAccessToken(uint(id))
}

func (s *Service) Logout(refreshToken string) error {
	_, err := s.jwt.ValidateToken(refreshToken, "refresh")
	return err
}
