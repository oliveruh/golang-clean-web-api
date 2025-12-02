package middleware

import (
	"net/http"
	"strings"

	"golang-clean-web-api/api/helper"
	"golang-clean-web-api/config"
	"golang-clean-web-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// Authentication middleware
func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := jwt.NewTokenService(cfg)

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helper.GenerateBaseResponse(nil, false, helper.AuthError))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helper.GenerateBaseResponse(nil, false, helper.AuthError))
			return
		}

		token := parts[1]
		claims, err := tokenService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helper.GenerateBaseResponse(nil, false, helper.AuthError))
			return
		}

		// Store user info in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
