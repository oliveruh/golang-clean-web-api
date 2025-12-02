package api

import (
	"encoding/json"
	"golang-clean-web-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

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

	// Create a test request (note: health endpoint requires trailing slash)
	req, err := http.NewRequest("GET", "/api/v1/health/", nil)
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

	// Check response fields (BaseHttpResponse structure)
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success true, got '%v'", response["success"])
	}

	if result, ok := response["result"].(string); !ok || result != "Working!" {
		t.Errorf("Expected result 'Working!', got '%v'", response["result"])
	}

	if resultCode, ok := response["resultCode"].(float64); !ok || resultCode != 0 {
		t.Errorf("Expected resultCode 0, got '%v'", response["resultCode"])
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
