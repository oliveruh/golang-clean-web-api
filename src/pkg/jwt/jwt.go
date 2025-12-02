package jwt

import (
	"errors"
	"time"

	"golang-clean-web-api/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type TokenService struct {
	config *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{config: cfg}
}

// GenerateAccessToken generates a new access token
func (s *TokenService) GenerateAccessToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(s.config.Jwt.AccessExpireTime * time.Minute)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Jwt.Secret))
}

// GenerateRefreshToken generates a new refresh token
func (s *TokenService) GenerateRefreshToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(s.config.Jwt.RefreshExpireTime * time.Minute)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Jwt.Secret))
}

// ValidateToken validates the token and returns the claims
func (s *TokenService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
