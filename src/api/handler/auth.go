package handler

import (
	"net/http"

	"golang-clean-web-api/api/dto"
	"golang-clean-web-api/api/helper"
	"golang-clean-web-api/config"
	"golang-clean-web-api/domain/model"
	"golang-clean-web-api/infra/persistence/database"
	"golang-clean-web-api/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	config       *config.Config
	tokenService *jwt.TokenService
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		config:       cfg,
		tokenService: jwt.NewTokenService(cfg),
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration information"
// @Success 201 {object} helper.BaseHttpResponse "User registered successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Validation error"
// @Failure 409 {object} helper.BaseHttpResponse "User already exists"
// @Failure 500 {object} helper.BaseHttpResponse "Internal server error"
// @Router /v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	db := database.GetDb()

	// Check if user exists
	var existingUser model.User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	// Create user
	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		IsActive: true,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	c.JSON(http.StatusCreated,
		helper.GenerateBaseResponse(gin.H{"user_id": user.Id}, true, helper.Success))
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.TokenResponse} "Login successful"
// @Failure 400 {object} helper.BaseHttpResponse "Validation error"
// @Failure 401 {object} helper.BaseHttpResponse "Invalid credentials"
// @Failure 500 {object} helper.BaseHttpResponse "Internal server error"
// @Router /v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	db := database.GetDb()

	// Find user
	var user model.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized,
			helper.GenerateBaseResponse(nil, false, helper.AuthError))
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized,
			helper.GenerateBaseResponse(nil, false, helper.AuthError))
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized,
			helper.GenerateBaseResponse(nil, false, helper.AuthError))
		return
	}

	// Generate tokens
	accessToken, err := h.tokenService.GenerateAccessToken(uint(user.Id), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	refreshToken, err := h.tokenService.GenerateRefreshToken(uint(user.Id), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	response := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK,
		helper.GenerateBaseResponse(response, true, helper.Success))
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Get a new access token using a refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.TokenResponse} "Token refreshed successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Validation error"
// @Failure 401 {object} helper.BaseHttpResponse "Invalid token"
// @Failure 500 {object} helper.BaseHttpResponse "Internal server error"
// @Router /v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// Validate refresh token
	claims, err := h.tokenService.ValidateToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			helper.GenerateBaseResponse(nil, false, helper.AuthError))
		return
	}

	// Generate new tokens
	accessToken, err := h.tokenService.GenerateAccessToken(claims.UserID, claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	refreshToken, err := h.tokenService.GenerateRefreshToken(claims.UserID, claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
		return
	}

	response := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK,
		helper.GenerateBaseResponse(response, true, helper.Success))
}
