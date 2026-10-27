package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	accessTTL          time.Duration
	refreshTTL         time.Duration
}
type Claims struct {
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func NewJWTManager(
	accessTokenSecret string,
	refreshTokenSecret string,
	acessTTL time.Duration,
	refreshTTL time.Duration,
) *JWTManager {
	return &JWTManager{
		accessTokenSecret:  []byte(accessTokenSecret),
		refreshTokenSecret: []byte(refreshTokenSecret),
		accessTTL:          acessTTL,
		refreshTTL:         refreshTTL,
	}
}

func (m *JWTManager) GenerateAccessToken(userID uint) (string, error) {
	return m.generateToken(userID, "access")
}
func (m *JWTManager) GenerateRefreshToken(userID uint) (string, error) {
	return m.generateToken(userID, "refresh")
}

func (m *JWTManager) generateToken(userID uint, tokenType string) (string, error) {
	now := time.Now()
	var secret []byte
	var ttl time.Duration
	if tokenType == "access" {
		secret = m.accessTokenSecret
		ttl = m.accessTTL
	}

	if tokenType == "refresh" {
		secret = m.refreshTokenSecret
		ttl = m.refreshTTL
	}

	claims := Claims{
		TokenType: tokenType,
		Subject:   strconv.FormatUint(uint64(userID), 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(secret)
}

func (m *JWTManager) ValidateToken(tokenString string, expectedType string) (*Claims, error) {
	var secret []byte
	switch expectedType {
	case "access":
		secret = m.accessTokenSecret
	case "refresh":
		secret = m.refreshTokenSecret
	default:
		return nil, ErrInvalidToken
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil || !token.Valid || claims.TokenType != expectedType {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
