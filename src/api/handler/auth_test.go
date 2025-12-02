package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-clean-web-api/api/dto"
	"golang-clean-web-api/config"

	"github.com/gin-gonic/gin"
)

func TestAuthHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()

	// Note: This test requires a database connection
	// In a real scenario, you would mock the database
	t.Skip("Skipping test that requires database connection")

	handler := NewAuthHandler(cfg)

	router := gin.Default()
	router.POST("/register", handler.Register)

	registerReq := dto.RegisterRequest{
		Username: "testuser",
		Password: "password123",
		Email:    "test@example.com",
	}

	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusConflict {
		t.Errorf("Expected status 201 or 409, got %d", w.Code)
	}
}

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()

	// Note: This test requires a database connection
	// In a real scenario, you would mock the database
	t.Skip("Skipping test that requires database connection")

	handler := NewAuthHandler(cfg)

	router := gin.Default()
	router.POST("/login", handler.Login)

	loginReq := dto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return 401 if user doesn't exist or credentials are wrong
	if w.Code != http.StatusUnauthorized && w.Code != http.StatusOK {
		t.Errorf("Expected status 200 or 401, got %d", w.Code)
	}
}
