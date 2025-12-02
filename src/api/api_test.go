package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"golang-clean-web-api/config"

	"github.com/gin-gonic/gin"
)

func TestHealthEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Load environment configuration
	cfg := config.GetConfig()

	// Create a test router
	r := gin.New()
	r.Use(gin.Recovery())

	// Register routes
	RegisterRoutes(r, cfg)

	// Create a test request
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse response
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Check response fields
	if status, ok := response["status"].(string); !ok || status != "ok" {
		t.Errorf("Expected status 'ok', got '%v'", response["status"])
	}

	if message, ok := response["message"].(string); !ok || message != "Server is running" {
		t.Errorf("Expected message 'Server is running', got '%v'", response["message"])
	}
}

func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())

	// Create test group
	testHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "ok"})
	}

	api := r.Group("/test")
	api.GET("/cors", testHandler)

	// Test GET request
	req, _ := http.NewRequest("GET", "/test/cors", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check response is OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
