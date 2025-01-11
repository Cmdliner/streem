package handler

import (
	"fmt"
	"net/http"

	"github.com/Cmdliner/streem/internal/config"
	dto "github.com/Cmdliner/streem/internal/dtos"
	"github.com/Cmdliner/streem/internal/model"
	"github.com/Cmdliner/streem/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	registeredUser, err := h.Service.Register(&user)
	if err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"user": registeredUser})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	authToken, err := h.Service.Login(&service.UserLogin{Email: user.Email, Password: user.Password})

	if err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return 
	}
	c.SetCookie("authCookie", authToken, 90000, "/", cfg.Server.URI, false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "User login successful", "token": authToken})
}

func (h *AuthHandler) ForgotPassword(c * gin.Context) {
	var email string
	c.BindJSON(&email)
	code, err := h.Service.ForgotPassword(email)
	if err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "An OTP code has been sent to your email", "code": code})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var data dto.PasswordReset
	c.BindJSON(&data)

	_, err := h.Service.UpdatePassword(data.Email, data.Code, data.Password)
	if err != nil {
		fmt.Print(fmt.Errorf("error: %w", err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
