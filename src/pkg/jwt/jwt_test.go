package jwt

import (
	"testing"

	"golang-clean-web-api/config"
)

func TestTokenService_GenerateAndValidateToken(t *testing.T) {
	// Create a test config
	cfg := &config.Config{
		Jwt: config.JwtConfig{
			Secret:            "test-secret-key-for-jwt-testing-purposes",
			AccessExpireTime:  60,
			RefreshExpireTime: 10080,
		},
	}

	service := NewTokenService(cfg)

	// Test data
	userID := uint(123)
	username := "testuser"

	// Test access token generation
	t.Run("Generate Access Token", func(t *testing.T) {
		token, err := service.GenerateAccessToken(userID, username)
		if err != nil {
			t.Fatalf("Failed to generate access token: %v", err)
		}
		if token == "" {
			t.Fatal("Generated token is empty")
		}
	})

	// Test refresh token generation
	t.Run("Generate Refresh Token", func(t *testing.T) {
		token, err := service.GenerateRefreshToken(userID, username)
		if err != nil {
			t.Fatalf("Failed to generate refresh token: %v", err)
		}
		if token == "" {
			t.Fatal("Generated token is empty")
		}
	})

	// Test token validation
	t.Run("Validate Token", func(t *testing.T) {
		token, err := service.GenerateAccessToken(userID, username)
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}

		claims, err := service.ValidateToken(token)
		if err != nil {
			t.Fatalf("Failed to validate token: %v", err)
		}

		if claims.UserID != userID {
			t.Errorf("Expected user ID %d, got %d", userID, claims.UserID)
		}
		if claims.Username != username {
			t.Errorf("Expected username %s, got %s", username, claims.Username)
		}
	})

	// Test invalid token
	t.Run("Validate Invalid Token", func(t *testing.T) {
		_, err := service.ValidateToken("invalid.token.here")
		if err == nil {
			t.Fatal("Expected error for invalid token, got nil")
		}
	})

	// Test token with wrong secret
	t.Run("Validate Token With Wrong Secret", func(t *testing.T) {
		// Generate token with one secret
		token, err := service.GenerateAccessToken(userID, username)
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}

		// Try to validate with different secret
		wrongSecretCfg := &config.Config{
			Jwt: config.JwtConfig{
				Secret:            "different-secret-key",
				AccessExpireTime:  60,
				RefreshExpireTime: 10080,
			},
		}
		wrongSecretService := NewTokenService(wrongSecretCfg)

		_, err = wrongSecretService.ValidateToken(token)
		if err == nil {
			t.Fatal("Expected error when validating token with wrong secret, got nil")
		}
	})
}
