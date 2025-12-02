package middleware

import (
	"time"

	"golang-clean-web-api/config"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

// RateLimiter creates a rate limiting middleware
func RateLimiter(cfg *config.Config) gin.HandlerFunc {
	if !cfg.RateLimiter.Enabled {
		// If rate limiting is disabled, return a no-op middleware
		return func(c *gin.Context) {
			c.Next()
		}
	}

	// Create a new rate limiter
	lmt := tollbooth.NewLimiter(float64(cfg.RateLimiter.RequestsPerMin)/60.0, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Minute,
	})

	// Set custom message
	lmt.SetMessage("Rate limit exceeded. Please try again later.")

	return tollbooth_gin.LimitHandler(lmt)
}
