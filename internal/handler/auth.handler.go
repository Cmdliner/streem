package handler

import (
	"net/http"

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

func (r *AuthHandler) Register(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	registeredUser, err := r.Service.Register(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"user": registeredUser})
}

func (r *AuthHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	authToken, err := r.Service.Login(&service.UserLogin{Email: user.Email, Password: user.Password})

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	c.SetCookie("authCookie", authToken, 90000, "/", "http://localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "User login successful", "token": authToken})
}

